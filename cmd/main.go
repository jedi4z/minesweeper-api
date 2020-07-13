package main

import (
	"github.com/jedi4z/minesweeper-api/app/adapters/rest_adapter"
	"github.com/jedi4z/minesweeper-api/app/container"
	"github.com/jedi4z/minesweeper-api/app/repositories/mysql_repositories"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := mysql_repositories.NewMySQLClient()
	if err != nil {
		log.Panicf("failed to connect to mysql: %v", err)
	}
	defer db.Close()

	c := container.InitializeContainer(db)
	r := rest_adapter.NewRestEngine(c)

	if err := r.Run(":80"); err != nil {
		log.Panicf("failed to initialize the interface engine: %v", err)
	}
}
