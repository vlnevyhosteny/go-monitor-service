package entities

import (
	"github.com/google/uuid"
	"github.com/vlnevyhosteny/go-monitor-service/domain/valueobjects"
)

type Monitoring struct {
	ID       uuid.UUID             `json:"id"`
	Schedule valueobjects.Schedule `json:"schedule"`
}
