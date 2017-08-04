package util

import (
	"github.com/jung-kurt/gofpdf"
)

func init() {
	gofpdf.SetDefaultCompression(false)
	gofpdf.SetDefaultCatalogSort(true)
	gofpdf.SetDefaultCreationDate(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
}

func createPDFfile() {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
}

func importFonts() {
	pdf.AddFont("Crimson Text", "", "CrimsonText-Roman.json")
	pdf.AddFont("Crimson Text Bold", "", "CrimsonText-Bold.json")
	pdf.AddFont("Andalus", "", "Andalus.json")
}

func PageHeader() {

}

func PageFooter() {

}

func GenerateReport()  {

}
