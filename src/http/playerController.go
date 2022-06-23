package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerController interface {
	GetPlayerByUsername(ctx *gin.Context)
	GetAllPlayers(ctx *gin.Context)
	CreatePlayer(ctx *gin.Context)
	UpdatePlayer(ctx *gin.Context)
	DeletePlayer(ctx *gin.Context)
}

type playerController struct {
	// TODO - implement
	// Service is here
}

func NewPlayerController() PlayerController {
	return &playerController{}
}

func (p *playerController) GetPlayerByUsername(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello GetPlayerByUsername!",
	})
}

func (p *playerController) GetAllPlayers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello GetAllPlayers!",
	})
}

func (p *playerController) CreatePlayer(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Hello CreatePlayer!",
	})
}

func (p *playerController) UpdatePlayer(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello UpdatePlayer!",
	})
}

func (p *playerController) DeletePlayer(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Hello DeletePlayer!",
	})
}
