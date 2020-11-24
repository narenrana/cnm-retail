package models

type AuthLoginRequest struct {
	Email        string `json:"email,omitempty"`
	Password     string   `json:"password,omitempty"`
	RememberMe   bool   `json:"rememberMe,omitempty"`
}
