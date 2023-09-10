package endpointmonitoring

import (
	"database/sql"

	"github.com/vlnevyhosteny/go-monitor-service/domain/aggregates"
)

type EndpointMonitoringSQLRepository struct {
	db *sql.DB
}

func NewEndpointMonitoringSQLRepository(db *sql.DB) *EndpointMonitoringSQLRepository {
	return &EndpointMonitoringSQLRepository{db: db}
}

func (r *EndpointMonitoringSQLRepository) Add(entity aggregates.EndpointMonitoring) error {
	_, err := r.db.Exec("") // TODO: implement
	return err
}
