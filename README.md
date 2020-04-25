# pdfgen
Small tool to generate sample pdfs

## Usage

Build:
```
make build
```

Test:
```
make test
```

Install to use as a global binary:
```
make install
```

Usage:
```
./pdfgen
Usage of ./pdfgen:
  -f string
    	Paper format
  -fontsize float
    	Font size. Default is calculated based on the size of the format
  -h float
    	Height. Ignored if format is specified
  -o string
    	Orientation P for portrait, L for landscape. Ignored if height of width are specified (default "P")
  -output string
    	Output path (default "result.pdf")
  -p int
    	Page count (default 1)
  -text string
    	Text in the pdf (default "Test pdf")
  -w float
    	Width. Ignored if format is specified
```

## Examples

Landscape A5 with 100 pages
```
./pdfgen -f A5 -o L -p 100
```

Custom pdf 222x111 with 1 page with 222x111 text
```
./pdfgen -w 222 -h 111 -text "222x111"
```