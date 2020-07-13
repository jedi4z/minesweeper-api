package models

type User struct {
	CommonFields
	Email    string  `json:"email" sql:"type:varchar(255);not null" binding:"required,email"`
	Password string  `json:"password,omitempty" sql:"type:varchar(255);not null" binding:"required"`
	Games    []*Game `json:"-" sql:"foreignkey:UserID"`
}

func (u User) SanitizeUser() User {
	return User{
		CommonFields: CommonFields{
			ID:        u.CommonFields.ID,
			CreatedAt: u.CommonFields.CreatedAt,
			UpdatedAt: u.CommonFields.UpdatedAt,
		},
		Email: u.Email,
	}
}
