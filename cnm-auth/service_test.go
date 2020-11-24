package auth

import (
	"reflect"
	"shopping-cart/cnm-auth/services"
	"testing"
)

func TestNewService(t *testing.T) {
	tests := []struct {
		name string
		want services.Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := services.NewService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_login(t *testing.T) {
	type args struct {
		request AuthLoginRequest
	}
	tests := []struct {
		name    string
		args    args
		want    AuthLoginResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &services.service{}
			got, err := s.login(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_logout(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    AuthLogoutResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &services.service{}
			got, err := s.logout(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("logout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("logout() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_recoverPassword(t *testing.T) {
	type args struct {
		request AuthRecoverPasswordRequest
	}
	tests := []struct {
		name    string
		args    args
		want    AuthRecoverPasswordResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &services.service{}
			got, err := s.recoverPassword(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("recoverPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recoverPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_refreshToken(t *testing.T) {
	type args struct {
		request AuthRefreshTokenRequest
	}
	tests := []struct {
		name    string
		args    args
		want    AuthRefreshTokenResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &services.service{}
			got, err := s.refreshToken(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("refreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("refreshToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_signUp(t *testing.T) {
	type args struct {
		request AuthSignUpRequest
	}
	tests := []struct {
		name    string
		args    args
		want    AuthSignUpResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &services.service{}
			got, err := s.signUp(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("signUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("signUp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
