package models

import "time"

type CommonFields struct {
	ID        uint      `json:"id" sql:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
