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

func (s *loggingService) login(request authLoginRequest) (authLoginResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", "took", time.Since(begin))
	}(time.Now())
	return s.Service.login(request)
}


func (s *loggingService) logout(token string) (authLogoutResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", token, "took", time.Since(begin))
	}(time.Now())
	return s.Service.logout(token)
}

func (s *loggingService) recoverPassword(request authRecoverPasswordRequest) (authRecoverPasswordResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", "took", time.Since(begin))
	}(time.Now())
	return s.Service.recoverPassword(request)
}



func (s *loggingService)  refreshToken(request authRefreshTokenRequest) (authRefreshTokenResponse, error){
	panic("implement me")
}

func (s *loggingService)  signUp(request  authSignUpRequest) (authSignUpResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", "took", time.Since(begin))
	}(time.Now())
	return s.Service.signUp(request)
}