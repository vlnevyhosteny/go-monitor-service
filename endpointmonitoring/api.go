package endpointmonitoring

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/vlnevyhosteny/go-monitor-service/database"
	"github.com/vlnevyhosteny/go-monitor-service/server"
)

type (
	PostEndpointMonitoringBody struct {
		Name     string `json:"name" validate:"required"`
		URL      string `json:"url" validate:"required"`
		Interval int    `json:"interval" validate:"required"`
		Monthly  bool   `json:"monthly"`
		Weekly   bool   `json:"weekly"`
		Daily    bool   `json:"daily"`
		Hourly   bool   `json:"hourly"`
		Minutely bool   `json:"minutely"`
	}
)

var PostEndpointMonitoringController server.Controller = server.Controller{
	Method: server.POST,
	Path:   "/endpoint-monitoring",
	Body:   reflect.TypeOf(PostEndpointMonitoringBody{}),
	Handler: func(c echo.Context) error {
		service := Wire(database.Config.DatabaseFile)

		err := service.Add("testik")
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, "Hello, World!")
	},
}
