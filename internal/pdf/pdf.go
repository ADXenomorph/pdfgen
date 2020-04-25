package pdf

import (
	"fmt"
	"math"

	"github.com/jung-kurt/gofpdf"
)

type Params struct {
	gofpdf.SizeType
	FontSize   float64
	PageCount  int
	OutputPath string
	TextPrefix string
}

type Pdf struct {
	params *Params
	fpdf   *gofpdf.Fpdf
}

func NewPdf(p *Params) *Pdf {
	pdf := &Pdf{
		params: p,
		fpdf:   gofpdf.NewCustom(&gofpdf.InitType{Size: p.SizeType}),
	}

	pdf.CreatePdf()

	return pdf
}

func (pdf *Pdf) CreatePdf() {
	orientation := pdf.getOrientation()
	pdf.fpdf.SetMargins(0, 0, 0)
	pdf.fpdf.SetAutoPageBreak(true, 0)

	pdf.fpdf.SetFont("Arial", "", pdf.getFontSize())

	for i := 1; i <= pdf.params.PageCount; i++ {
		pdf.fpdf.AddPage()
		if orientation == gofpdf.OrientationLandscape {
			pdf.writeText(i)
		} else {
			pdf.writeVerticalText(i)
		}
	}
}

func (pdf *Pdf) OutputFile() error {
	return pdf.fpdf.OutputFileAndClose(pdf.params.OutputPath)
}

func (pdf *Pdf) getFontSize() float64 {
	if pdf.params.FontSize > 0 {
		return pdf.params.FontSize
	}

	shortestSide := math.Min(pdf.params.SizeType.Wd, pdf.params.SizeType.Ht)

	return shortestSide * 0.4
}

func (pdf *Pdf) getOrientation() string {
	var o string
	if pdf.params.SizeType.Wd > pdf.params.SizeType.Ht {
		o = gofpdf.OrientationLandscape
	} else {
		o = gofpdf.OrientationPortrait
	}
	return o
}

func (pdf *Pdf) writeVerticalText(i int) {
	pdf.fpdf.TransformBegin()
	pdf.fpdf.TransformRotate(-90, pdf.params.SizeType.Wd/2, pdf.params.SizeType.Ht/2)
	pdf.writeText(i)
	pdf.fpdf.TransformEnd()
}

func (pdf *Pdf) writeText(i int) {
	pdf.fpdf.CellFormat(
		pdf.params.SizeType.Wd,
		pdf.params.SizeType.Ht,
		fmt.Sprintf("%s page %d", pdf.params.TextPrefix, i),
		"",
		0,
		"CM",
		false,
		0,
		"",
	)
}
