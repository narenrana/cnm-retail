package auth

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) login(user string, password string, rememberMe string) (authLoginResponse, error){
	defer func(begin time.Time) {
		s.requestCount.With("method", "track").Add(1)
		s.requestLatency.With("method", "track").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.login(user,password,rememberMe)
}
func (s *instrumentingService) logout(token string) (authLogoutResponse, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "track").Add(1)
		s.requestLatency.With("method", "track").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.logout(token)
}
