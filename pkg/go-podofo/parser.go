package podofo

import (
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/denisss025/go-podofo/internal/pdf"
	"github.com/denisss025/go-podofo/internal/set"
)

const (
	versionLen          = 3
	magicLen            = 8
	xrefEntrySize       = 20
	xrefBuf             = 512
	maxXRefSessionCount = 512
)

type Reader interface {
	io.ReadSeeker
}

// Parser reads a PDF file into memory. The file
// can be modifierd and written back using
// Writer. Most PDF features are supported.
type Parser struct {
	visitedXRefOffsets *set.Set[int64]

	tokenizer *Tokenizer
	entries   []XRefEntry
	objects   *IndirectObjectList
	trailer   *ParserObject
	encrypt   *Encrypt

	password string

	magicOffset   int64
	fileSize      int64
	xrefOffset    int64
	lastEOFOffset int64

	incUpdatesCount int

	pdfVersion PDFVersion

	loadOnDemand  bool
	strictParsing bool
	ignoreBroken  bool
	hasXrefStream bool
}

// Parse opens a PDF file and parses it.
//
// If loadOnDemand is true all objects will be read
// from the file at the time they are accessed first. If false all the objects will be read immidiately.
// This is faster if you do not need the complete PDF
// file in memory.
func Parse(
	r Reader, options ...ParserOption,
) (*Parser, error) {
	p := &Parser{
		tokenizer:          NewTokenizer(),
		objects:            new(IndirectObjectList),
		trailer:            new(ParserObject),
		visitedXRefOffsets: set.New[int64](),
	}

	p.reset()

	for _, opt := range options {
		opt(p)
	}

	ok, err := p.isPDFFile(r)
	if err != nil {
		return nil, fmt.Errorf("pdf parse: %w", err)
	}

	if !ok {
		return nil, fmt.Errorf("pdf parse: %w", ErrNoPDFFile)
	}

	if err = p.readDocumentStructure(r); err != nil {
		return nil, fmt.Errorf("pdf parse: %w", err)
	}

	if err = p.readObjects(r); err != nil {
		p = nil
		err = fmt.Errorf("pdf parse: %w", err)
	}

	return p, err
}

// IsEncrypted returns true if the PDF file is encrypted.
func (p *Parser) IsEncrypted() bool { return p.encrypt != nil }

// TakeEncrypt gives the encryption object from the
// parser. The internal handle will be set to nil and
// the ownership of the object is given to the caller.
//
// Only call this if yout need access to the encryption object.
func (p *Parser) TakeEncrypt() *Encrypt {
	encrypt := p.encrypt
	p.encrypt = nil

	return encrypt
}

// GetEncrypt gives the encryption object from the parser.
func (p *Parser) GetEncrypt() *Encrypt { return p.encrypt }

// Trailer returns the trailer object of the PDF file.
func (p *Parser) Trailer() Object { return p.trailer }

// Password retuns the PDF file password.
func (p *Parser) Password() string { return p.password }

// NumIncrementalUpdates returns the number of incremental
// updates that have been applied to the last parsed
// PDF file.
//
// The function returs 0 if there is no updade has been
// applied.
func (p *Parser) NumIncrementalUpdates() int {
	return p.incUpdatesCount
}

// Objects returns a reference to the sorted internal
// objects list.
func (p *Parser) Objects() *IndirectObjectList {
	return p.objects
}

// PDFVersion returns the file format version of the PDF.
func (p *Parser) PDFVersion() PDFVersion {
	return p.pdfVersion
}

// LoadOnDemand returns true if the LoadOnDemand option
// was applied for Parser.
func (p *Parser) LoadOnDemand() bool { return p.loadOnDemand }

// FileSize returns the length of the file in bytes.
func (p *Parser) FileSize() int64 { return p.fileSize }

// IsStrictParsing returns true if strict parsing mode
// is enabled.
func (p *Parser) IsStrictParsing() bool { return p.strictParsing }

// XRefOffset returns the offset to the XRef stream.
func (p *Parser) XRefOffset() int64 { return p.xrefOffset }

// HasXRefStream returns true if the file has XRef stream.
func (p *Parser) HasXRefStream() bool { return p.hasXrefStream }

func (p *Parser) readDocumentStructure(r Reader) (err error) {
	p.fileSize, err = r.Seek(0, io.SeekEnd)
	if err != nil {
		return fmt.Errorf("read document structure: %w", err)
	}

	if err = p.checkEOFMarker(r); err != nil {
		return fmt.Errorf("read document structure: %w", err)
	}

	if p.xrefOffset, err = p.findXRef(r); err != nil {
		return fmt.Errorf("read document structure: %w", err)
	}

	if err = p.readXRefContents(r, p.xrefOffset, false); err != nil {
		return fmt.Errorf("read document structure: %w", err)
	}

	if p.trailer != nil && p.trailer.IsDictionary() {
		entriesCount := FindKey[int64](p.trailer.Object.(*Dictionary), NameKeySize, -1)

		if entriesCount >= 0 && len(p.entries) > int(entriesCount) {
			// Total number of xref entries to read is greater than the /Size
			// specified in the trailer if any. That's an error unless we're
			// trying to recover from a missing /Size entry.
			log.Printf("There are more objects %v in this XRef table than secified in the size key of the trailer directory (%v)!",
				len(p.entries), entriesCount)
		}
	}

	return nil
}

func (p *Parser) readXRefContents(r Reader, offset int64, atEnd bool) (err error) {
	var firstObject, objectCount int64

	if p.visitedXRefOffsets.Contains(offset) {
		return fmt.Errorf("read xref contents: %w: cycle in xref structure: offset %d already visited", ErrInvalidXRef, offset)
	}

	p.visitedXRefOffsets.Put(offset)

	currentPosition, _ := r.Seek(0, io.SeekCurrent)
	fileSize, _ := r.Seek(0, io.SeekEnd)
	_, _ = r.Seek(currentPosition, io.SeekStart)

	if offset > fileSize {
		if _, err = p.findXRef(r); err != nil {
			return fmt.Errorf("read xref contents: %w", err)
		}

		offset, _ = r.Seek(0, io.SeekCurrent)
		p.tokenizer.buffer = pdf.ResizeSlice(p.tokenizer.buffer, xrefBuf*4)

		err = p.findTokenBackward(r, "xref", int64(len(p.tokenizer.buffer)), offset)
		if err != nil {
			return fmt.Errorf("read xref contents: %w", err)
		}

		p.tokenizer.buffer = p.tokenizer.buffer[:xrefBuf]
		offset, _ = r.Seek(0, io.SeekCurrent)
		p.xrefOffset = offset

	} else {
		_, _ = r.Seek(offset, io.SeekStart)
	}

	token, err := p.tokenizer.TryReadNextToken(r)
	if err != nil {
		return fmt.Errorf("read xref contents: read next token: %w", ErrNoXRef)
	}

	if string(token) != "xref" {
		if p.pdfVersion < PDFVersion13 {
			return fmt.Errorf("read xref contents: %w", ErrNoXRef)
		}

		p.hasXrefStream = true

		if err = p.readXRefStreamContents(r, offset, atEnd); err != nil {
			err = fmt.Errorf("read xref contents: %w", err)
		}

		return err
	}

	var xrefSectionCount int

	for {
		if xrefSectionCount == maxXRefSessionCount {
			return fmt.Errorf("read xref contents: %w", ErrNoEOFToken)
		}

		token, err := p.tokenizer.TryReadNextToken(r)
		if err != nil {
			return fmt.Errorf("read xref contents: %w", ErrNoXRef)
		}

		if string(token) == "trailer" {
			break
		}

		firstObject, err = NewTokenizer().ReadNextNumber(r)
		if err != nil {
			if errors.Is(err, ErrNoNumber) ||
				errors.Is(err, ErrInvalidXRef) ||
				errors.Is(err, ErrUnexpectedEOF) {
				break
			}

			return fmt.Errorf("read xref contents: %w", err)
		}

		objectCount, err = NewTokenizer().ReadNextNumber(r)
		if err != nil {
			if errors.Is(err, ErrNoNumber) ||
				errors.Is(err, ErrInvalidXRef) ||
				errors.Is(err, ErrUnexpectedEOF) {
				break
			}

			return fmt.Errorf("read xref contents: %w", err)
		}

		if atEnd {
			_, err = r.Seek(objectCount*xrefEntrySize, io.SeekStart)
			if err != nil {
				return fmt.Errorf("read xref contents: %w", err)
			}
		} else {
			err = p.readXRefSubsection(r, firstObject, objectCount)
			if err != nil {
				if errors.Is(err, ErrNoNumber) ||
					errors.Is(err, ErrInvalidXRef) ||
					errors.Is(err, ErrUnexpectedEOF) {
					break
				}

				return fmt.Errorf("read xref contents: %w", err)
			}
		}
	}

	if err = p.readNextTrailer(r); err != nil {
		if !errors.Is(err, ErrNoTrailer) {
			return fmt.Errorf("read xref contents: %w", err)
		}
	}

	return nil
}

func (p *Parser) readXRefSubsection(r Reader, firstObject int64, objectCount int64) (err error) {
	panic("not implemented") // TODO: implement me
}

func (p *Parser) readXRefStreamContents(r Reader, offser int64, trailerOnly bool) error {
	panic("not implemented") // TODO: implement me
}

func (p *Parser) readNextTrailer(r Reader) error {
	panic("not implemented") // TODO: implement me
}

func (p *Parser) readObjects(r Reader) (err error) {
	if p.trailer == nil {
		return fmt.Errorf("read objects: %w", ErrNoTrailer)
	}

	encrypt := p.trailer.Dictionary().Key(NameKeyEncrypt)
	if encrypt != nil && encrypt.Kind() != ObjectKindNull {
		p.encrypt, err = EncryptFromObject(encrypt)
		if err != nil {
			return fmt.Errorf("read objects: %w", err)
		}
	}

	if err = p.readObjectsInternal(r); err != nil {
		err = fmt.Errorf("read objects: %w", err)
	}

	return err
}

func (p *Parser) isPDFFile(r Reader) (ok bool, err error) {
	const pdfStart = "%PDF-"

	_, err = r.Seek(0, io.SeekStart)
	if err != nil {
		return ok, fmt.Errorf("read pdf header: %w", err)
	}

	buf := make([]byte, BufferSize)

	n, err := r.Read(buf[:len(pdfStart)])
	if err != nil {
		return ok, fmt.Errorf("read pdf header: %w", err)
	}

	if n != len(pdfStart) || pdfStart != string(buf[:len(pdfStart)]) {
		return false, nil
	}

	n, err = r.Read(buf[:versionLen])
	if err != nil {
		return ok, fmt.Errorf("read pdf header: verion: %w", err)
	}

	if n != versionLen {
		return false, nil
	}

	p.magicOffset, _ = r.Seek(0, io.SeekCurrent)
	p.pdfVersion = getPDFVersion(string(buf[:versionLen]))

	return p.pdfVersion == PDFVersionUnknown, nil
}

func (p *Parser) findTokenBackward(r Reader, token string, bytesRange int64, searchEnd int64) error {
	panic("not implemented") // TODO: implement me
}

func (p *Parser) mergeTrailer(trailer Object) error {
	panic("not implemented") // TODO: implement me
}

func (p *Parser) findXRef(r Reader) (offset int64, err error) {
	panic("not implemented") // TODO: implement me
}

func (p *Parser) readObjectsInternal(r Reader) error {
	compressedObjects := make(map[int][]int)

	for i := range p.entries {
		entry := &p.entries[i]

		if entry.Parsed {
			switch entry.Type {
			case XRefEntryTypeInUse:
				if entry.Offset > 0 {
					ref := NewReference(i, entry.Generation)
					obj, err := NewParserObject(r, entry.Offset,
						WithDocument(p.objects.Document()), WithReference(ref),
						DelayedLoad(p.loadOnDemand))
					if err != nil {
						// TODO: ignore broken objects
						return fmt.Errorf("read object internal: %w", err)
					}

					obj.Encrypt = *p.encrypt

					if p.encrypt != nil && obj.IsDictionary() {
						typeObj := obj.Object.(*Dictionary).Key(NameKeyType)

						name, ok := typeObj.(Name)
						if ok && name == NameXRef {
							// TODO: ignore broken objects
							// XRef is never encrypted
							panic("not implemented") // TODO: implement me
						}
					}

					p.objects.PushObject(obj)
				} else if entry.Generation == 0 {
					if p.strictParsing {
						return fmt.Errorf("read object internal: %w: found object with 0 offset which should be 'f' instead of 'n'", ErrInvalidXRef)
					}

					// TODO? log?

					p.objects.AddFreeObject(NewReference(i, FirstGeneration))
				}
			case XRefEntryTypeFree:
				if i > 0 {
					p.objects.SafeAddFreeObject(NewReference(i, entry.Generation))
				}
			case XRefEntryTypeCompressed:
				compressedObjects[entry.ObjectNum] = append(compressedObjects[entry.ObjectNum], i)
			default:
				return fmt.Errorf("read object internal: %w", ErrInvalidEnumValue)
			}
		} else if i > 0 { // unparsed
			p.objects.AddFreeObject(NewReference(i, FirstGeneration))
		}
		// the linked free list in the xref section is not always correct in pdf's
		// (especially Illustrator) but Acrobat still accepts them. I've seen XRefs
		// where some object-numbers are alltogether missing and multiple XRefs where
		// the link list is broken.
		// Because PdfIndirectObjectList relies on a unbroken range, fill the free list more
		// robustly from all places which are either free or unparsed
	}

	// all normal objects including object streams are available now,
	// we can parse the object streams safely now.
	//
	// Note that even if demand loading is enabled we still currently read all
	// objects from the stream into memory then free the stream.
	//

	for objNo, list := range compressedObjects {
		p.readCompressedObjectFromStream(objNo, list)
	}

	if !p.loadOnDemand {
		// Force loading of streams. We can't do this during the initial
		// run that populates m_Objects because a stream might have a /Length
		// key that references an object we haven't yet read. So we must do it here
		// in a second pass, or (if demand loading is enabled) defer it for later.
		panic("not implemented") // TODO: implement me
	}

	if err := p.updateDocumentVersion(); err != nil {
		return fmt.Errorf("read object internal: %w", err)
	}

	return nil
}

func (p *Parser) readCompressedObjectFromStream(objNo int, objects []int) error {
	panic("not implemented") // TODO: implement me
}

func (p *Parser) checkEOFMarker(r Reader) error {
	panic("not implemented") // TODO: implement me
}

func (p *Parser) reset() {
	p.pdfVersion = DefaultPDFVersion
	p.loadOnDemand = false
	p.magicOffset = 0
	p.hasXrefStream = false
	p.xrefOffset = 0
	p.lastEOFOffset = 0
	p.trailer = nil
	p.entries = p.entries[:0]
	p.encrypt = nil
	p.ignoreBroken = true
	p.incUpdatesCount = 0
}

func (p *Parser) documentID() String {
	panic("not implemented") // TODO: implement me
}

func (p *Parser) updateDocumentVersion() error {
	panic("not implemented") // TODO: implement me
}

func (p *Parser) parseEncrypt(r Reader, obj Object) (err error) {
	switch obj := obj.(type) {
	case *Reference:
		i := obj.NumObjects()
		if i <= 0 || i >= len(p.entries) {
			return fmt.Errorf("parse encrypt object: %w", ErrInvalidEncryptionDict)
		}

		parserObj, err := NewParserObject(r, p.entries[i].Offset)
		if err != nil {
			return fmt.Errorf("parse encrypt object: %w", err)
		}

		if err = parserObj.Parse(); err != nil {
			return fmt.Errorf("parse encrypt object: %w", err)
		}

		p.entries[i].Parsed = false

		p.encrypt, err = EncryptFromObject(parserObj)
		if err != nil {
			return fmt.Errorf("parse encrypt object: %w", err)
		}
	case *Dictionary:
		p.encrypt, err = EncryptFromObject(obj)
		if err != nil {
			return fmt.Errorf("parse encrypt dict: %w", err)
		}
	default:
		return fmt.Errorf("parse encrypt: %w", ErrInvalidEncryptionDict)
	}

	isAuthenticated := p.encrypt.Authenticate(p.password, p.documentID())
	if !isAuthenticated {
		return fmt.Errorf("parse encrypt: auth: %w", ErrInvalidPassword)
	}

	return err
}

func getPDFVersion(version string) (ver PDFVersion) {
	switch version {
	case "1.0":
		ver = PDFVersion10
	case "1.1":
		ver = PDFVersion11
	case "1.2":
		ver = PDFVersion12
	case "1.3":
		ver = PDFVersion13
	case "1.4":
		ver = PDFVersion14
	case "1.5":
		ver = PDFVersion15
	case "1.6":
		ver = PDFVersion16
	case "1.7":
		ver = PDFVersion17
	case "2.0":
		ver = PDFVersion20
	}

	return ver
}
