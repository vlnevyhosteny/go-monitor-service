package endpointmonitoring

import "errors"

var (
	ErrFailedToAddEndpointMonitoring = errors.New("failed to add endpoint monitoring")
)

type EndpointMonitoringRepository interface {
	Add(EndpointMonitoring) error
}
