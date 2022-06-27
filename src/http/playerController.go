package http

import (
	"net/http"

	"github.com/felipemaxplay/first-go-api/src/http/data/request"
	"github.com/felipemaxplay/first-go-api/src/http/data/response"
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
	result, err := p.playerService.GetPlayer(username)
	if err != nil {
		res := response.BuildPlayerError(http.StatusNotFound, err.Error(), "Player not found.")
		ctx.JSON(res.Code, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (p *playerController) GetAllPlayers(ctx *gin.Context) {
	players, err := p.playerService.GetAllPlayers()
	if err != nil {
		res := response.BuildPlayerError(http.StatusBadRequest, err.Error(), "Failed to process request.")
		ctx.JSON(res.Code, res)
	}

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

	result, err := p.playerService.CreatePlayer(playerDto)
	if err != nil {
		res := response.BuildPlayerError(http.StatusUnprocessableEntity, err.Error(), "Failed to processe request.")
		ctx.JSON(res.Code, res)
	}

	ctx.JSON(http.StatusCreated, result)
}

func (p *playerController) UpdatePlayer(ctx *gin.Context) {
	var playerDto request.PlayerRequestDto
	username := ctx.Param("username")
	err := ctx.ShouldBind(&playerDto)
	if err != nil {
		res := response.BuildPlayerError(http.StatusBadRequest, err.Error(), "Failed to process request")
		ctx.JSON(res.Code, res)
		return
	}

	result, err := p.playerService.UpdatePlayer(username, playerDto)
	if err != nil {
		res := response.BuildPlayerError(http.StatusUnprocessableEntity, err.Error(), "Failed to process request")
		ctx.JSON(res.Code, res)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (p *playerController) DeletePlayer(ctx *gin.Context) {
	username := ctx.Param("username")
	err := p.playerService.DeletePlayer(username)
	if err != nil {
		res := response.BuildPlayerError(http.StatusNotFound, err.Error(), "Player not found.")
		ctx.JSON(res.Code, res)
		return
	}

	ctx.JSON(http.StatusNoContent, "")
}
