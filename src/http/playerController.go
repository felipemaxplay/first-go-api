package http

import (
	"fmt"
	"net/http"

	"github.com/felipemaxplay/first-go-api/src/http/data/request"
	"github.com/felipemaxplay/first-go-api/src/http/data/response"
	"github.com/felipemaxplay/first-go-api/src/model"
	"github.com/felipemaxplay/first-go-api/src/service"
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
	playerService service.PlayerService
}

func NewPlayerController(service service.PlayerService) PlayerController {
	return &playerController{
		playerService: service,
	}
}

func (p *playerController) GetPlayerByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	fmt.Println(username)
	var result model.Player = p.playerService.GetPlayer(username)
	if (result == model.Player{}) {
		res := response.BuildPlayerError(http.StatusBadRequest, "", "")
		ctx.JSON(res.Code, res)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (p *playerController) GetAllPlayers(ctx *gin.Context) {
	players := p.playerService.GetAllPlayers()
	ctx.JSON(http.StatusOK, players)
}

func (p *playerController) CreatePlayer(ctx *gin.Context) {
	var playerDto request.PlayerRequestDto
	errDto := ctx.ShouldBind(&playerDto)
	if errDto != nil {
		res := response.BuildPlayerError(http.StatusBadRequest, errDto.Error(), "Failed to processe request")
		ctx.JSON(res.Code, res)
		return
	}
	result := p.playerService.CreatePlayer(playerDto)
	ctx.JSON(http.StatusCreated, result)
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
