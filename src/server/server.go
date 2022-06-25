package server

import (
	"github.com/felipemaxplay/first-go-api/src/config"
	"github.com/felipemaxplay/first-go-api/src/http"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	controller http.PlayerController
}

func NewServer(controller http.PlayerController) Server {
	return &server{
		controller: controller,
	}
}

func (s *server) Run() {
	cfg := config.GetServer()
	hostPort := cfg.Host + ":" + cfg.Port
	r := gin.Default()

	routes := r.Group("/api/v1/players")
	{
		routes.GET("/:username", s.controller.GetPlayerByUsername)
		routes.GET("", s.controller.GetAllPlayers)
		routes.POST("", s.controller.CreatePlayer)
		routes.PUT("/:username", s.controller.UpdatePlayer)
		routes.DELETE("/:username", s.controller.DeletePlayer)
	}

	r.Run(hostPort)
}
