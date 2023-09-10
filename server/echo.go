package server

import (
	"net/http"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/vlnevyhosteny/go-monitor-service/database"
)

func AddController(server *echo.Echo, controller EchoController) {
	handler := func(c echo.Context) error {
		defer database.ProvideConnection(database.Config.DatabaseFile).Close()

		if controller.ValidateBody() {
			u := reflect.PointerTo(controller.GetBodyType())

			if err := c.Bind(u); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			if err := c.Validate(u); err != nil {
				return err
			}
		}

		return controller.GetHandler()(c)
	}

	switch controller.GetMethod() {
	case GET:
		server.GET(controller.GetPath(), handler)
	case POST:
		server.POST(controller.GetPath(), handler)
	case PUT:
		server.PUT(controller.GetPath(), handler)
	case DELETE:
		server.DELETE(controller.GetPath(), handler)
	}
}

func AddControllers(server *echo.Echo, controllers []EchoController) {
	for _, controller := range controllers {
		AddController(server, controller)
	}
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (validator *CustomValidator) Validate(i interface{}) error {
	if err := validator.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
