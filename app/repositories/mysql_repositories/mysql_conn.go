package mysql_repositories

import (
	"github.com/jedi4z/minesweeper-api/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

func NewMySQLClient() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "/tmp/minesweeper.db")
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(
		&models.User{},
		&models.Game{},
		&models.Row{},
		&models.Cell{},
	)

	return db, nil
}
