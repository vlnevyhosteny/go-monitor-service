package endpointmonitoring

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vlnevyhosteny/go-monitor-service/server"
)

var PostEndpointMonitoringController server.Controller = server.Controller{
	Method: server.POST,
	Path:   "/endpoint-monitoring",
	Handler: func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	},
}
