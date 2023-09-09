package server

import "github.com/labstack/echo/v4"

func AddController(server *echo.Echo, controller EchoController) {
	switch controller.GetMethod() {
	case GET:
		server.GET(controller.GetPath(), controller.GetHandler())
	case POST:
		server.POST(controller.GetPath(), controller.GetHandler())
	case PUT:
		server.PUT(controller.GetPath(), controller.GetHandler())
	case DELETE:
		server.DELETE(controller.GetPath(), controller.GetHandler())
	}
}

func AddControllers(server *echo.Echo, controllers []EchoController) {
	for _, controller := range controllers {
		AddController(server, controller)
	}
}
