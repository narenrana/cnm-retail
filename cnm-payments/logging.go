package payments

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) add(request  paymentsAddRequest) (paymentsAddResponse , error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", request, "took", time.Since(begin))
	}(time.Now())
	return s.Service.add(request)
}


func (s *loggingService) list() ( paymentsListResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", "took", time.Since(begin))
	}(time.Now())
	return s.Service.list()
}