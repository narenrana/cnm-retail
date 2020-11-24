package auth

import (
	"shopping-cart/cnm-auth/models"
	"shopping-cart/cnm-auth/services"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	services.Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s services.Service) services.Service {
	return &loggingService{logger, s}
}

func (s *loggingService) login(request models.AuthLoginRequest) (models.AuthLoginResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", "took", time.Since(begin))
	}(time.Now())
	return s.Service.Login(request)
}


func (s *loggingService) logout(token string) (models.AuthLogoutResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", token, "took", time.Since(begin))
	}(time.Now())
	return s.Service.Logout(token)
}

func (s *loggingService) recoverPassword(request models.AuthRecoverPasswordRequest) (models.AuthRecoverPasswordResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", "took", time.Since(begin))
	}(time.Now())
	return s.Service.RecoverPassword(request)
}



func (s *loggingService)  refreshToken(request models.AuthRefreshTokenRequest) (models.AuthRefreshTokenResponse, error){
	panic("implement me")
}

func (s *loggingService)  signUp(request models.AuthSignUpRequest) (models.AuthSignUpResponse, error) {
	defer func(begin time.Time) {
		s.logger.Log("method", "track", "tracking_id", "took", time.Since(begin))
	}(time.Now())
	return s.Service.SignUp(request)
}