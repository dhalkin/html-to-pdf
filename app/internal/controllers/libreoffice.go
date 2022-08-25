package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// /usr/bin/libreoffice --headless --convert-to pdf /tmp/example1.html --outdir /tmp
// /usr/bin/soffice --headless --convert-to pdf /tmp/example1.html --outdir /tmp

func TransferLibreoffice(c echo.Context) error {
	var reqBody []byte
	reqBody, _ = ioutil.ReadAll(c.Request().Body)

	tempFile, err := ioutil.TempFile("/tmp", "tempo*.html")
	if err != nil {
		return err
	}

	if _, err := tempFile.WriteString(string(reqBody)); err != nil {
		tempFile.Close()
		return err
	}
	tempFile.Close()

	app := "libreoffice"
	arg0 := "--headless"
	arg1 := "--convert-to"
	arg2 := "pdf:writer_web_pdf_Export"
	arg3 := tempFile.Name()
	arg4 := "--outdir"
	arg5 := "/tmp"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5)
	_, err = cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	newPath2 := strings.Replace(tempFile.Name(), "html", "pdf", -1)

	defer func() {
		os.Remove(tempFile.Name())
		os.Remove(newPath2)
	}()

	return c.File(newPath2)
}
