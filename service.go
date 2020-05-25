package samwise

import (
	"github.com/krystalmejia24/samwise/db"

	log "github.com/rs/zerolog"
)

// Service is the implemenation of Samwise svc
type Service struct {
	DBConn db.Svc
	Logger log.Logger
}

// NewSvc creates and returns a samwise service
func NewSvc(db db.Svc, logger log.Logger) *Service {
	return &Service{
		db,
		logger,
	}
}

// Healthy returns the health of the service
func (Service) Healthy() bool {
	return true
}
