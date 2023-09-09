package entities

import (
	"github.com/google/uuid"
	"github.com/vlnevyhosteny/go-monitor-service/domain/valueobjects"
)

type ResultStatus int

const (
	Success ResultStatus = iota
	Failure
	Pending
)

type Result struct {
	ID       uuid.UUID             `json:"id"`
	Response valueobjects.Response `json:"response"`
	Status   ResultStatus          `json:"status"`
}
