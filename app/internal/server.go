package internal

import (
	"github.com/dhalkin/html-to-pdf/app/internal/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
)

const appPort = ":8080"

func StartServer() {
	e := getE()
	if err := e.Start(appPort); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func getE() *echo.Echo {
	e := echo.New()
	e.HideBanner = false
	e.Debug = true
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.BodyLimit("1M"))

	// add middleware and routes
	e.POST("/transfer-wkhtmltopdf", controllers.TransferWkhtmltopdf)
	e.POST("/transfer-libreoffice", controllers.TransferLibreoffice)

	return e
}
