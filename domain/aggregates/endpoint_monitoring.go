package aggregates

import (
	"errors"

	"github.com/google/uuid"
	"github.com/vlnevyhosteny/go-monitor-service/domain/entities"
	"github.com/vlnevyhosteny/go-monitor-service/domain/valueobjects"
)

var (
	ErrEmptyEndpointName = errors.New("empty endpoint name")
)

type (
	EndpointMonitoring struct {
		endpoint    *entities.Endpoint
		monitorings []*entities.Monitoring
		results     []*entities.Result
	}

	EndpointMonitoringRepository interface {
		Add(EndpointMonitoring) error
	}

	EndpointMonitoringService interface {
		Add(name string) error
	}
)

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

func (e EndpointMonitoring) Endpoint() *entities.Endpoint {
	return e.endpoint
}

func (e EndpointMonitoring) Monitorings() []*entities.Monitoring {
	return e.monitorings
}

func (e EndpointMonitoring) Results() []*entities.Result {
	return e.results
}
