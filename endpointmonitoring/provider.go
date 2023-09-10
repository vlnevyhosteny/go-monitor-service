package endpointmonitoring

import (
	"database/sql"
	"sync"

	"github.com/google/wire"
	"github.com/vlnevyhosteny/go-monitor-service/database"
	"github.com/vlnevyhosteny/go-monitor-service/domain/aggregates"
)

var (
	service     *Service
	serviceOnce sync.Once

	repository     *EndpointMonitoringSQLRepository
	repositoryOnce sync.Once

	Set wire.ProviderSet = wire.NewSet(
		ProvideRepository,
		ProvideService,
		database.ProvideConnection,

		wire.Bind(new(aggregates.EndpointMonitoringRepository), new(*EndpointMonitoringSQLRepository)),
		wire.Bind(new(aggregates.EndpointMonitoringService), new(*Service)),
	)
)

func ProvideService(repo aggregates.EndpointMonitoringRepository) *Service {
	serviceOnce.Do(func() {
		service = NewEndpointMonitoringService(repo)
	})

	return service
}

func ProvideRepository(db *sql.DB) *EndpointMonitoringSQLRepository {
	repositoryOnce.Do(func() {
		repository = NewEndpointMonitoringSQLRepository(db)
	})

	return repository
}
