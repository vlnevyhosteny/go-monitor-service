package endpointmonitoring

import (
	"github.com/vlnevyhosteny/go-monitor-service/domain/aggregates"
	"github.com/vlnevyhosteny/go-monitor-service/domain/valueobjects"
)

type Service struct {
	endpointMonitorings aggregates.EndpointMonitoringRepository
}

func NewEndpointMonitoringService(repository aggregates.EndpointMonitoringRepository) *Service {
	return &Service{endpointMonitorings: repository}
}

func (s *Service) Add(name string) error {
	endpointMonitoring, err := aggregates.NewEndpointMonitoring(name, valueobjects.Address{}) // todo: implement address
	if err != nil {
		return err
	}

	return s.endpointMonitorings.Add(endpointMonitoring)
}
