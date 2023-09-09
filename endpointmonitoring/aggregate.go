package endpointmonitoring

import (
	"errors"

	"github.com/google/uuid"
	"github.com/vlnevyhosteny/go-monitor-service/domain/entities"
	"github.com/vlnevyhosteny/go-monitor-service/domain/valueobjects"
)

var (
	ErrEmptyEndpointName = errors.New("empty endpoint name")
)

type EndpointMonitoring struct {
	endpoint    *entities.Endpoint
	monitorings []*entities.Monitoring
	results     []*entities.Result
}

func NewEndpointMonitoring(name string, address valueobjects.Address) (EndpointMonitoring, error) {
	if name == "" {
		return EndpointMonitoring{}, ErrEmptyEndpointName
	}

	return EndpointMonitoring{
		endpoint: &entities.Endpoint{
			ID:      uuid.New(),
			Name:    name,
			Address: address,
		},
		monitorings: make([]*entities.Monitoring, 0),
		results:     make([]*entities.Result, 0),
	}, nil
}
