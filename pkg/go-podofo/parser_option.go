package podofo

// ParserOption is used for PDF file Parser.
type ParserOption func(*Parser)

// LoadOnDemand tells Parser to read object at the time
// they are accessed first.
//
// This is faster if you do not need the complete PDF
// file in memory.
func LoadOnDemand() ParserOption {
	return func(p *Parser) { p.loadOnDemand = true }
}

// StrictMode enables strict parsing mode.
//
// If you enable strict parsing, Parse will fail
// on a few more common PDF failuers.
// Please note that Parser is by default very strict
// already and does not recover from e.g. wrong XREF
// tables.
func StrictMode() ParserOption {
	return func(p *Parser) { p.strictParsing = true }
}

// DontIgnoreBrokenObjects tells the parser to not ignore broken
// object, i.e. XRef entries that do not point to valid
// objects.
func DontIgnoreBrokenObjects() ParserOption {
	return func(p *Parser) { p.ignoreBroken = true }
}

// Password sets the PDF file password.
func Password(password string) ParserOption {
	return func(p *Parser) { p.password = password }
}
