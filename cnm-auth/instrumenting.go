package auth

import (
	"shopping-cart/cnm-auth/models"
	"shopping-cart/cnm-auth/services"
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

func (s *instrumentingService) login(request models.AuthLoginRequest) (models.AuthLoginResponse, error){
	defer func(begin time.Time) {
		s.requestCount.With("method", "track").Add(1)
		s.requestLatency.With("method", "track").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Login(request)
}
func (s *instrumentingService) logout(token string) (models.AuthLogoutResponse, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "track").Add(1)
		s.requestLatency.With("method", "track").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Logout(token)
}



func (s *instrumentingService) recoverPassword(request models.AuthRecoverPasswordRequest) (models.AuthRecoverPasswordResponse, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "track").Add(1)
		s.requestLatency.With("method", "track").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.RecoverPassword(request)
}



func (s *instrumentingService)  refreshToken(request models.AuthRefreshTokenRequest) (models.AuthRefreshTokenResponse, error){
	defer func(begin time.Time) {
		s.requestCount.With("method", "track").Add(1)
		s.requestLatency.With("method", "track").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.RefreshToken(request)
}

func (s *instrumentingService)  signUp(request models.AuthSignUpRequest) (models.AuthSignUpResponse, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "track").Add(1)
		s.requestLatency.With("method", "track").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.SignUp(request)
}