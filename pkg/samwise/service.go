package samwise

import (
	"github.com/krystalmejia24/samwise/pkg/db"

	log "github.com/sirupsen/logrus"
)

// Service is the implemenation of Samwise svc
type Service struct {
	DB     *db.Repo
	Logger *log.Logger
}

// NewSvc creates and returns a samwise service
func NewSvc(db *db.Repo, logger *log.Logger) *Service {
	return &Service{
		db,
		logger,
	}
}

// Healthy returns the health of the service
func (Service) Healthy() bool {
	return true
}
