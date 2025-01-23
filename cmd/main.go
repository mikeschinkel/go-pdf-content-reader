package main

import (
	"fmt"
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
