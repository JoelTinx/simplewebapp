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

}

func importFonts() {

}

func PageHeader() {

}

func PageFooter() {

}

func GenerateReport()  {

}
