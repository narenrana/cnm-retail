package offers

import (
	"time"

	"github.com/go-kit/kit/log"
)

type offersLoggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &offersLoggingService{logger, s}
}

func (s *offersLoggingService) add(request  addOffersRequest) (addOffersResponse , error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", request, "took", time.Since(begin))
	}(time.Now())
	return s.Service.add(request)
}


func (s *offersLoggingService) list() ( getOffersResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", "took", time.Since(begin))
	}(time.Now())
	return s.Service.list()
}