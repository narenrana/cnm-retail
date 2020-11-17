package models

type AuthSignUpResponse struct {
	Email          string `json:"email,omitempty"`
	Token          string  `json:"token,omitempty"`
	Err        error  `json:"error,omitempty"`

}
