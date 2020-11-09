package carts

import (
	"github.com/go-kit/kit/log"
	"shopping-cart/cnm-carts/models"
	"shopping-cart/cnm-carts/services"
	"time"
)

type loggingService struct {
	logger log.Logger
	services.Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s services.Service) services.Service {
	return &loggingService{logger, s}
}

func (s *loggingService) add(request models.AddToCartRequest) (models.CartResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", request, "took", time.Since(begin))
	}(time.Now())
	return s.Service.Add(request)
}


func (s *loggingService) list(request models.GetCartRequest) (models.CartResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", "took", time.Since(begin))
	}(time.Now())
	return s.Service.Get(request)
}