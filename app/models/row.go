package models

type Row struct {
	CommonFields
	GameID uint    `json:"-"`
}
