package podofo

import "github.com/denisss025/go-podofo/internal/pdf"

type Metadata struct {
	Author   string
	Creator  string
	Title    string
	Subject  string
	Keywords []string
}

type MemDocument struct {
	base     *pdf.Document
	Metadata Metadata
}

type DocumentOptionFunc func(doc *MemDocument)

func WithCreator(name string) DocumentOptionFunc {
	return func(doc *MemDocument) { doc.Metadata.Creator = name }
}

func WithTitle(title string) DocumentOptionFunc {
	return func(doc *MemDocument) { doc.Metadata.Title = title }
}

func WithAuthor(name string) DocumentOptionFunc {
	return func(doc *MemDocument) { doc.Metadata.Author = name }
}

func WithSubjects(subj string) DocumentOptionFunc {
	return func(doc *MemDocument) { doc.Metadata.Subject = subj }
}

func WithKeyword(keywords ...string) DocumentOptionFunc {
	return func(doc *MemDocument) {
		doc.Metadata.Keywords = append(
			doc.Metadata.Keywords,
			keywords...,
		)
	}
}

func NewMemDocument(options ...DocumentOptionFunc) *MemDocument {
	doc := &MemDocument{base: pdf.NewDocument()}

	for _, option := range options {
		option(doc)
	}

	return doc
}

func (doc *MemDocument) AddPage(size PageSize) *Page {
	panic("not implemented") // TODO: implement me
}

func (doc *MemDocument) FindFont(name string) *Font {
	panic("not implemented") // TODO: implement me
}

func (doc *MemDocument) SaveToFile(filename string) error {
	panic("not implemented") // TODO: implement me
}
