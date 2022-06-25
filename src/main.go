package main

import (
	"log"

	"github.com/felipemaxplay/first-go-api/src/config"
	"github.com/felipemaxplay/first-go-api/src/database"
	"github.com/felipemaxplay/first-go-api/src/http"
	"github.com/felipemaxplay/first-go-api/src/repository"
	"github.com/felipemaxplay/first-go-api/src/server"
	"github.com/felipemaxplay/first-go-api/src/service"
	"gorm.io/gorm"
)

var (
	errConfig     error                       = config.LoadConfig()
	db            *gorm.DB                    = database.LoadDB()
	playerRepo    repository.PlayerRepository = repository.NewPlayerRepository(db)
	playerService service.PlayerService       = service.NewPlayerService(playerRepo)
	controller    http.PlayerController       = http.NewPlayerController(playerService)
	sv            server.Server               = server.NewServer(controller)
)

func main() {
	if errConfig != nil {
		log.Println(errConfig)
	}
	sv.Run()
}
