# PDF Reader

This is a quick-and-dirty proof-of-concept to read text content from a PDF file, including the 
line breaks.  See the repo I forked this from for more info on the original sources. 

## Read plain text
This is the `main.go` in the `/cmd` directory. To try it run `make`. 

```golang
package main

import (
	"fmt"
	"io"
	"log/slog"
	"strings"

	pdf "github.com/mikeschinkel/go-pdf-content-reader"
)

const filename = "test.pdf"

func main() {
	var err error
	var pdfRdr *pdf.Reader
	//var ioRdr io.Reader
	var lines []string
	var sb strings.Builder

	pdfRdr, err = pdf.Open(filename)
	// remember close file
	defer mustClose(pdfRdr)
	if err != nil {
		goto end
	}

	lines, err = pdfRdr.GetPlainTextLines()
	if err != nil {
		goto end
	}
	for i, line := range lines {
		sb.WriteString(fmt.Sprintf("%d. — %s\n", i, strings.TrimRight(line, "\n")))
	}
	fmt.Print(sb.String())
end:
	if err != nil {
		slog.Error(err.Error())
	}
}

func mustClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		slog.Error("ERROR: Failed to close", "error", "err")
	}
}```
