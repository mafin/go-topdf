package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
)

func main() {
	file := "test.txt"
	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("%s file not found", file)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 14)

	pdf.MultiCell(190, 5, string(content), "0", "0", false)

	_ = pdf.OutputFileAndClose("test.pdf")

	fmt.Println("PDF Created")
}
