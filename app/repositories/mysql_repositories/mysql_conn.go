package mysql_repositories

import (
	"fmt"
	"github.com/jedi4z/minesweeper-api/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewMySQLClient() (*gorm.DB, error) {
	connectionUri := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"minesweeper_user",
		"qwerty",
		"database",
		"minesweeper",
	)

	db, err := gorm.Open("mysql", connectionUri)
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(
		&models.Game{},
		&models.Row{},
		&models.Cell{},
	)

	return db, nil
}
