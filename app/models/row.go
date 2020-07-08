package models

type Row struct {
	CommonFields
	Cells  []*Cell `json:"cells" sql:"foreignkey:RowID"`
	GameID uint    `json:"-"`
}
