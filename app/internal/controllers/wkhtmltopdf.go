package controllers

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"os"
)

func TransferWkhtmltopdf(c echo.Context) error {
	var reqBody []byte
	reqBody, _ = ioutil.ReadAll(c.Request().Body)

	tempFile, err := ioutil.TempFile("/tmp", "html2pdf*.html")
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(string(reqBody)); err != nil {
		tempFile.Close()
		return err
	}
	tempFile.Close()
	// strange bug have to reopen the file again
	tempFile, err = os.Open(tempFile.Name())
	if err != nil {
		return err
	}

	// init wkhtmltopdf
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	page := wkhtmltopdf.NewPageReader(tempFile)
	pdfg.AddPage(page)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	pdfFile, err := ioutil.TempFile("/tmp", "html2pdf*.pdf")
	if err != nil {
		return err
	}
	defer os.Remove(pdfFile.Name())

	err = pdfg.WriteFile(pdfFile.Name())
	if err != nil {
		return err
	}

	if err := pdfFile.Close(); err != nil {
		return err
	}

	return c.File(pdfFile.Name())
}
