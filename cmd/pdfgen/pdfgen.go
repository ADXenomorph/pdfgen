package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"

	"github.com/ADXenomorph/pdfgen/internal/pdf"
)

func main() {
	format := flag.String("f", "", "Paper format")
	bleed := flag.Bool("b", false, "Add an 8 mm bleed area to the size. Ignored if height or width are specified")
	orientation := flag.String("o", "P", "Orientation P for portrait, L for landscape. Ignored if height or width are specified")

	width := flag.Float64("w", 0, "Width. Ignored if format is specified")
	height := flag.Float64("h", 0, "Height. Ignored if format is specified")

	pages := flag.Int("p", 1, "Page count")

	path := flag.String("output", "result.pdf", "Output path")
	prefix := flag.String("text", "Test pdf", "Text in the pdf")
	font := flag.Float64("fontsize", 0, "Font size. Default is calculated based on the size of the format")

	flag.Parse()

	if *format == "" && (*width == 0 || *height == 0) {
		fmt.Print("You need to specify either a format or width+height\n\n")
		flag.Usage()
		os.Exit(1)
	}

	var size *gofpdf.SizeType
	if *format != "" {
		var found bool
		size, found = pdf.GetSize(*format, *orientation, *bleed)
		if !found {
			fmt.Printf("Unknown format %s\n", *format)
			fmt.Println(pdf.GetFormatsHelp())
			os.Exit(1)
		}
	} else {
		size = &gofpdf.SizeType{Wd: *width, Ht: *height}
	}

	params := &pdf.Params{
		SizeType:   *size,
		FontSize:   *font,
		PageCount:  *pages,
		OutputPath: *path,
		TextPrefix: *prefix,
	}

	pdf := pdf.NewPdf(params)
	err := pdf.OutputFile()

	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Saved to %s\n", params.OutputPath)
}
