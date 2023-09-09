package entities

import (
	"github.com/google/uuid"
	"github.com/vlnevyhosteny/go-monitor-service/domain/valueobjects"
)

type Endpoint struct {
	ID   uuid.UUID        `json:"id"`
	Name string           `json:"name"`
	Url  valueobjects.Url `json:"url"`
}
