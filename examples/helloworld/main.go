package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	podofo "github.com/denisss025/go-podofo/pkg/go-podofo"
)

// TODO: make example work

func printHelp() {
	fmt.Println("This is an example for the Go-PoDoFo library.")
	fmt.Println("It creates a small PDF file containing the",
		"text >Hello World!<")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Printf("  %s [outputfile.pdf]\n\n", path.Base(os.Args[0]))
}

func helloWorld(filename string) error {
	doc := podofo.NewMemDocument(
		podofo.WithCreator(
			"example-helloworld - A Go-PoDoFo test application",
		),
		podofo.WithAuthor("Denis Novikov"),
		podofo.WithTitle("Hello World"),
		podofo.WithSubjects("Testing the Go-PoDoFo Library"),
		podofo.WithKeyword("Test", "PDF", "Hello World"),
	)

	// TODO: add page
	page := doc.AddPage(podofo.PageSizeA4())

	// TODO: painter ctor
	painter := podofo.NewPainter()
	painter.SetCanvas(page)

	// TODO: search fonts
	font := doc.FindFont("Arial")
	if font == nil {
		return fmt.Errorf("%w: cannot find font", podofo.ErrInvalidHandle)
	}

	// TODO: font metrics
	metrics := font.Metrics()
	log.Println("The font name is", metrics.FontName())
	log.Println("The family font name is", metrics.FontFamilyName())
	log.Println("The font file path is", metrics.FilePath())
	log.Println("The font face index is", metrics.FaceIndex())

	// TODO: draw text
	err := painter.DrawText("ABCDEFGHIKLMNOPQRSTVXYZ", 56.69, page.Rect().Height-56.69)
	if err != nil {
		return fmt.Errorf("ascii: %w", err)
	}

	err = painter.DrawText("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЬЫЭЮЯ", 56.69, page.Rect().Height-80)
	if errors.Is(err, podofo.ErrInvalidFontData) {
		log.Printf("WARNING: The matched font %q does not support cyrillic.", metrics.FontName())
	} else if err != nil {
		return fmt.Errorf("cyrillic: %w", err)
	}

	// TODO: finish drawing
	err = painter.FinishDrawing()
	if err != nil {
		return fmt.Errorf("finish: %w", err)
	}

	// TODO: save to file (io.Reader)
	return doc.SaveToFile(filename)
}

func main() {
	if len(os.Args) != 2 {
		printHelp()
		os.Exit(1)
	}

	if err := helloWorld(os.Args[1]); err != nil {
		log.Fatal("Cannot create file: ", err)
	}

	log.Println("Created a PDF file containing the line",
		`"Hello World!":`, os.Args[1])
}
