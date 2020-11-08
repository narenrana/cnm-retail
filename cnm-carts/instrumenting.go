package carts

import (
	"shopping-cart/cnm-carts/models"
	"shopping-cart/cnm-carts/services"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	services.Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s services.Service) services.Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) add(request models.AddToCartRequest) (models.AddToCartResponse, error){
	defer func(begin time.Time) {
		s.requestCount.With("method", "track").Add(1)
		s.requestLatency.With("method", "track").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Add(request)
}

func (s *instrumentingService) list(request models.GetCartRequest) (models.GetCartResponse, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "track").Add(1)
		s.requestLatency.With("method", "track").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Get(request)
}
