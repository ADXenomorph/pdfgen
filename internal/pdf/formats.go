package pdf

import (
	"fmt"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

var Formats = map[string]gofpdf.SizeType{
	"100x70":            gofpdf.SizeType{Wd: 707.00, Ht: 1000.00},
	"36x90_standard":    gofpdf.SizeType{Wd: 914.40, Ht: 2286.00},
	"500x700":           gofpdf.SizeType{Wd: 500.00, Ht: 700.00},
	"50x70":             gofpdf.SizeType{Wd: 500.00, Ht: 707.00},
	"5R":                gofpdf.SizeType{Wd: 119.00, Ht: 170.00},
	"5R_CA":             gofpdf.SizeType{Wd: 127.00, Ht: 178.00},
	"700x1000":          gofpdf.SizeType{Wd: 700.00, Ht: 1000.00},
	"850x2000_one_time": gofpdf.SizeType{Wd: 850.00, Ht: 2000.00},
	"850x2000_standard": gofpdf.SizeType{Wd: 850.00, Ht: 2000.00},
	"92x210":            gofpdf.SizeType{Wd: 210.00, Ht: 296.00},
	"A0":                gofpdf.SizeType{Wd: 841.00, Ht: 1189.00},
	"A1":                gofpdf.SizeType{Wd: 594.00, Ht: 841.00},
	"A2":                gofpdf.SizeType{Wd: 420.00, Ht: 594.00},
	"A3":                gofpdf.SizeType{Wd: 297.00, Ht: 420.00},
	"A4":                gofpdf.SizeType{Wd: 210.00, Ht: 297.00},
	"A5":                gofpdf.SizeType{Wd: 148.00, Ht: 210.00},
	"A6":                gofpdf.SizeType{Wd: 105.00, Ht: 148.00},
	"A7":                gofpdf.SizeType{Wd: 125.40, Ht: 176.20},
	"A9":                gofpdf.SizeType{Wd: 128.53, Ht: 204.73},
	"ArchC":             gofpdf.SizeType{Wd: 457.20, Ht: 609.60},
	"ArchD":             gofpdf.SizeType{Wd: 609.60, Ht: 914.40},
	"B5":                gofpdf.SizeType{Wd: 176.00, Ht: 250.00},
	"BB":                gofpdf.SizeType{Wd: 50.00, Ht: 90.00},
	"BC":                gofpdf.SizeType{Wd: 55.00, Ht: 90.00},
	"BD":                gofpdf.SizeType{Wd: 55.00, Ht: 85.00},
	"BX":                gofpdf.SizeType{Wd: 50.81, Ht: 88.91},
	"C4":                gofpdf.SizeType{Wd: 221.00, Ht: 316.00},
	"DL":                gofpdf.SizeType{Wd: 99.00, Ht: 210.00},
	"DX":                gofpdf.SizeType{Wd: 215.90, Ht: 299.30},
	"EC4":               gofpdf.SizeType{Wd: 221.00, Ht: 316.00},
	"LG":                gofpdf.SizeType{Wd: 139.71, Ht: 215.91},
	"LT":                gofpdf.SizeType{Wd: 215.91, Ht: 279.41},
	"MA":                gofpdf.SizeType{Wd: 210.00, Ht: 279.00},
	"SM":                gofpdf.SizeType{Wd: 107.96, Ht: 139.71},
	"XL":                gofpdf.SizeType{Wd: 279.41, Ht: 431.81},
}

func GetSize(format string, orientation string) (*gofpdf.SizeType, bool) {
	size, found := Formats[format]

	if !found {
		return nil, false
	}

	if orientation == "L" {
		size = gofpdf.SizeType{Wd: size.Ht, Ht: size.Wd}
	}

	return &size, true
}

func GetFormatsHelp() string {
	res := []string{
		"Format name: width height",
	}
	for k := range Formats {
		format := Formats[k]
		res = append(res, fmt.Sprintf("%s: %.2f %.2f", k, format.Wd, format.Ht))
	}

	return strings.Join(res, "\n")
}
