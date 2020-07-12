package models

type User struct {
	CommonFields
	Email    string  `json:"email" sql:"type:varchar(255);not null"`
	Password string  `json:"-" sql:"type:varchar(255);not null"`
	Games    []*Game `json:"-" sql:"foreignkey:UserID"`
}
