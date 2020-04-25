package pdf_test

import (
	"testing"

	"github.com/ADXenomorph/pdfgen/internal/pdf"
	"github.com/jung-kurt/gofpdf"
)

func TestPdf(t *testing.T) {
	params := &pdf.Params{
		SizeType:   gofpdf.SizeType{Ht: 148, Wd: 105},
		FontSize:   14,
		PageCount:  10,
		OutputPath: "../../test.pdf",
		TextPrefix: "Random text",
	}

	pdf := pdf.NewPdf(params)
	pdf.OutputFile()
}
