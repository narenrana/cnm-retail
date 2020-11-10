package entities


type Users struct {
	UserId         *int        `gorm:"primaryKey" json:"userId,omitempty"`
	FirstName      string      `json:"firstName,omitempty"`
	MiddleName     string      `json:"middleName,omitempty"`
	LastName       string      `json:"lastName,omitempty"`
	UserEmail      string      `json:"userEmail,omitempty"`
	Password       string      `json:"password,omitempty"`
	PhoneNumber    string      `json:"phoneNumber,omitempty"`

}