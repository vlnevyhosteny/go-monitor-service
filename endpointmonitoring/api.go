package endpointmonitoring

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vlnevyhosteny/go-monitor-service/database"
	"github.com/vlnevyhosteny/go-monitor-service/server"
)

var PostEndpointMonitoringController server.Controller = server.Controller{
	Method: server.POST,
	Path:   "/endpoint-monitoring",
	Handler: func(c echo.Context) error {
		service := Wire(database.Config.DatabaseFile)
		defer database.ProvideConnection(database.Config.DatabaseFile).Close() // TODO: find more generic wat to close connection

		err := service.Add("testik")
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, "Hello, World!")
	},
}
