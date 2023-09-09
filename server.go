package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vlnevyhosteny/go-monitor-service/endpointmonitoring"
	"github.com/vlnevyhosteny/go-monitor-service/server"
)

var controllers = []server.EchoController{endpointmonitoring.PostEndpointMonitoringController}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	server.AddControllers(e, controllers)

	e.Logger.Fatal(e.Start(":3000"))
}
