package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vlnevyhosteny/go-monitor-service/database"
	"github.com/vlnevyhosteny/go-monitor-service/endpointmonitoring"
	"github.com/vlnevyhosteny/go-monitor-service/server"
)

var controllers = []server.EchoController{endpointmonitoring.PostEndpointMonitoringController}

func main() {
	e := echo.New()

	db := database.ProvideConnection(database.Config.DatabaseFile)
	err := database.Init(db)
	if err != nil {
		e.Logger.Fatal(err)
		panic(err)
	}

	e.Validator = &server.CustomValidator{Validator: validator.New()}

	e.Use(middleware.Logger())

	server.AddControllers(e, controllers)

	e.Logger.Fatal(e.Start(":3000"))
}
