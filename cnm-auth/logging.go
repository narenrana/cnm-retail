package auth

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

func (s *loggingService) login(user string, password string, rememberMe string) (authLoginResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", user, "took", time.Since(begin))
	}(time.Now())
	return s.Service.login(user,password,rememberMe)
}


func (s *loggingService) logout(token string) (authLogoutResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", token, "took", time.Since(begin))
	}(time.Now())
	return s.Service.logout(token)
}