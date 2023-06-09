package podofo

import (
	"bytes"
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
func (p *Parser) Objects() *IndirectObjectList { return p.objects }

// PDFVersion returns the file format version of the PDF.
func (p *Parser) PDFVersion() PDFVersion { return p.pdfVersion }

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

	if p.trailer != nil && p.trailer.Dictionary != nil {
		entriesCount := p.trailer.Dictionary.Int(pdf.KeySize, -1)

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
		if p.pdfVersion < pdf.Version13 {
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
	defer func() {
		if err != nil {
			err = fmt.Errorf("read xref subsection: %w", err)
		}
	}()

	if firstObject < 0 {
		return fmt.Errorf("%w: value object is negative", ErrValueOutOfRange)
	}

	if objectCount < 0 {
		return fmt.Errorf("%w: object count is negative", ErrValueOutOfRange)
	}

	p.entries = pdf.ResizeSlice(p.entries, int(firstObject+objectCount))

	// TODO?
	// consume all whitespaces
	//     char ch;
	//     while (device.Peek(ch) && m_tokenizer.IsWhitespace(ch))
	//         (void)device.ReadChar();

	panic("not implemented") // TODO: implement me
}

func (p *Parser) readXRefStreamContents(r Reader, offset int64, trailerOnly bool) error {
	_, err := r.Seek(offset, io.SeekStart)
	if err != nil {
		return fmt.Errorf("read xref stream contents: %w", err)
	}

	panic("not implemented") // TODO: implement me
}

func (p *Parser) readNextTrailer(r Reader) error {
	panic("not implemented") // TODO: implement me
}

func (p *Parser) readObjects(r Reader) (err error) {
	if p.trailer == nil || p.trailer.Dictionary == nil {
		return fmt.Errorf("read objects: %w", ErrNoTrailer)
	}

	encrypt := p.trailer.Dictionary.Key(pdf.KeyEncrypt)
	if encrypt != nil && encrypt.Kind() != pdf.ObjectKindNull {
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

	buf := make([]byte, len(pdfStart)+versionLen)

	n, err := r.Read(buf)
	if err != nil {
		return ok, fmt.Errorf("read pdf header: %w", err)
	}

	if n != len(buf) || !bytes.HasPrefix(buf, []byte(pdfStart)) {
		return false, nil
	}

	p.magicOffset, _ = r.Seek(0, io.SeekCurrent)
	p.pdfVersion = PDFVersion(buf[:versionLen])

	return p.pdfVersion.Validate() == nil, nil
}

func (p *Parser) findTokenBackward(r Reader, token string, bytesRange int64, searchEnd int64) error {
	currpos, err := r.Seek(searchEnd, io.SeekStart)
	if err != nil {
		return fmt.Errorf("find token backward: %w", err)
	}

	searchSize := currpos
	if currpos > bytesRange {
		searchSize = bytesRange
	}

	_, err = r.Seek(-searchSize, io.SeekCurrent)
	if err != nil {
		return fmt.Errorf("find token backward: %w", err)
	}

	buffer := p.tokenizer.buffer[:searchSize]
	// search backwards in the buffer in case the buffer contains null bytes
	// because it is right after a stream (can't use strstr for this reason)

	i := int64(bytes.LastIndex(buffer, []byte(token)))
	if i <= 0 {
		// TODO? if (i == 0) ???
		return fmt.Errorf("find token backwards: %w", ErrInternalLogic)
	}

	_, err = r.Seek(searchEnd-(searchSize-i), io.SeekStart)
	if err != nil {
		err = fmt.Errorf("read token backwards: %w", err)
	}

	return err
}

func (p *Parser) mergeTrailer(trailer *Dictionary) {
	if p.trailer == nil {
		// TODO: create trailer
	}

	for _, key := range []pdf.Name{
		pdf.KeySize,
		pdf.KeyRoot,
		pdf.KeyEncrypt,
		pdf.KeyInfo,
		pdf.KeyID,
	} {
		obj := trailer.Key(key)

		if p.trailer.Key(key) == nil {
			p.trailer.AddKey(key, obj)
		}
	}
}

func (p *Parser) findXRef(r Reader) (offset int64, err error) {
	err = p.findTokenBackward(r, "startxref", xrefBuf, p.
		lastEOFOffset)
	if err != nil {
		return -1, fmt.Errorf("find xref: %w", err)
	}

	token, err := p.tokenizer.TryReadNextToken(r)
	if err != nil || string(token) != "startxref" {
		if !p.strictParsing {
			err = p.findTokenBackward(r, "startref", xrefBuf, p.
				lastEOFOffset)
			if err != nil {
				return -1, fmt.Errorf("find xref: %w", err)
			}

			token, err = p.tokenizer.TryReadNextToken(r)
			if err != nil || string(token) != "startref" {
				return -1, errors.Join(fmt.Errorf("find xref: %w", ErrNoXRef), err)
			}
		} else {
			return -1, errors.Join(fmt.Errorf("find xref: %w", ErrNoXRef), err)
		}
	}

	offset, err = p.tokenizer.ReadNextNumber(r)
	if err != nil {
		err = fmt.Errorf("find xref: %w", err)
	}

	return offset, err
}

func (p *Parser) readObjectsInternal(r Reader) error {
	compressedObjects := make(map[int][]int)

	for i := range p.entries {
		entry := &p.entries[i]

		if entry.Parsed {
			switch entry := entry.Entry.(type) {
			case XRefEntryInUse:
				if entry.Offset > 0 {
					ref := pdf.NewReference(i, entry.Generation)
					obj, err := NewParserObject(r, entry.Offset,
						WithDocument(p.objects.Document()), WithReference(ref),
						DelayedLoad(p.loadOnDemand))
					if err != nil {
						// TODO: ignore broken objects
						return fmt.Errorf("read object internal: %w", err)
					}

					obj.Encrypt = *p.encrypt

					if p.encrypt != nil && obj.Dictionary != nil {
						typeObj := obj.Dictionary.Key(pdf.KeyType)

						name, ok := typeObj.(*pdf.NameObject)
						if ok && name.Name == pdf.NameXRef {
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

					p.objects.AddFreeObject(pdf.NewReference(i, pdf.FirstGeneration))
				}
			case XRefEntryFree:
				if i > 0 {
					p.objects.SafeAddFreeObject(pdf.NewReference(i, entry.Generation))
				}
			case XRefEntryCompressed:
				compressedObjects[entry.ObjectNumber] = append(compressedObjects[entry.ObjectNumber], i)
			default:
				return fmt.Errorf("read object internal: %w", ErrInvalidEnumValue)
			}
		} else if i > 0 { // unparsed
			p.objects.AddFreeObject(pdf.NewReference(i, pdf.FirstGeneration))
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
	const (
		EOFToken    = "%%EOF"
		EOFTokenLen = len(EOFToken)
	)

	p.lastEOFOffset = 0

	currentPos, err := r.Seek(-int64(EOFTokenLen), io.SeekEnd)
	if err != nil {
		return fmt.Errorf("check EOF: %w", err)
	}

	if p.IsStrictParsing() {
		buf := make([]byte, EOFTokenLen)

		n, err := r.Read(buf)
		if err != nil {
			return fmt.Errorf("check EOF: %w", err)
		}

		// TODO? error?
		if n != EOFTokenLen {
			return fmt.Errorf("check EOF: %w", ErrUnknown)
		}

		if !bytes.Equal(buf, []byte(EOFToken)) {
			return fmt.Errorf("check EOF: %w", ErrNoEOFToken)
		}
	} else {
		panic("not implemented") // TODO: implement me
	}

	p.lastEOFOffset = currentPos

	return nil
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

func (p *Parser) documentID() (*String, error) {
	id := p.trailer.Dictionary.Key(pdf.KeyID)
	if id == nil {
		return nil, fmt.Errorf("get document ID: not found in trailer: %w", ErrInvalidEncryptionDict)
	}

	array, ok := id.(*Array)
	if !ok {
		return nil, fmt.Errorf("get document ID: not an array: %w", ErrInvalidEncryptionDict)
	}

	// TODO? check the type of the first object?
	return array.objects[0].(*String), nil
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

		parserObj, err := NewParserObject(r, p.entries[i].Entry.(XRefEntryInUse).Offset)
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

	documentID, err := p.documentID()
	if err != nil {
		return fmt.Errorf("parse encrypt: %w", err)
	}

	isAuthenticated := p.encrypt.Authenticate(p.password, documentID)
	if !isAuthenticated {
		return fmt.Errorf("parse encrypt: auth: %w", ErrInvalidPassword)
	}

	return err
}
