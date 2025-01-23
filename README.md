# PDF Reader

A simple Go library which enables reading PDF files. 

This repository contains merged code from https://github.com/rsc/pdf and https://github.com/ledongthuc/pdf

Features
  - Get plain text content (without formatting)
  - Get Content (including all font and formatting information)

## Install:

`go get -u github.com/mikeschinkel/go-pdf-content-reader`


## Read plain text

```golang
package main

import (
	"bytes"
	"fmt"

	"github.com/mikeschinkel/go-pdf-content-reader"
)

func main() {
	pdf.DebugOn = true
	content, err := readPdf("test.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	return
}

func readPdf(path string) (string, error) {
	r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
   defer r.Close()
	var buf bytes.Buffer
   b, err := r.GetPlainText()
   if err != nil {
      return "", err
   }
   buf.ReadFrom(b)
	return buf.String(), nil
}
```

## Read all text with styles from PDF

```golang
func readPdf2(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		var lastTextStyle pdf.Text
		texts := p.Content().Text
		for _, text := range texts {
			if isSameSentence(text, lastTextStyle) {
				lastTextStyle.S = lastTextStyle.S + text.S
			} else {
				fmt.Printf("Font: %s, Font-size: %f, x: %f, y: %f, content: %s \n", lastTextStyle.Font, lastTextStyle.FontSize, lastTextStyle.X, lastTextStyle.Y, lastTextStyle.S)
				lastTextStyle = text
			}
		}
	}
	return "", nil
}
```


## Read text grouped by rows

```golang
package main

import (
	"fmt"
	"os"

	"github.com/mikeschinkel/go-pdf-content-reader"
)

func main() {
	content, err := readPdf(os.Args[1]) // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	return
}

func readPdf(path string) (string, error) {
	r, err := pdf.Open(path)
	defer func() {
      _ = r.Close()
	}()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		for _, row := range rows {
		    println(">>>> row: ", row.Position)
		    for _, word := range row.Content {
		        fmt.Println(word.S)
		    }
		}
	}
	return "", nil
}
```
