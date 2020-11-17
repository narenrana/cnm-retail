package models

type AuthSignUpRequest struct {
	FirstName      string   `json:"firstName,omitempty"`
	MiddleName     string   `json:"middleName,omitempty"`
	LastName       string     `json:"lastName,omitempty"`
	Email          string     `json:"email,omitempty"`
	Password       string      `json:"password,omitempty"`
	PhoneNumber    string   `json:"phoneNumber,omitempty"`
}

