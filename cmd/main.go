package main

import (
	"github.com/jedi4z/minesweeper-api/app/adapters/rest_adapter"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := rest_adapter.NewRestEngine()

	if err := r.Run(":8080"); err != nil {
		log.Panicf("failed to initialize the interface engine: %v", err)
	}
}
