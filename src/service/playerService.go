package service

import (
	"github.com/felipemaxplay/first-go-api/src/http/data/request"
	"github.com/felipemaxplay/first-go-api/src/model"
	"github.com/felipemaxplay/first-go-api/src/repository"
)

type PlayerService interface {
	CreatePlayer(playerDto request.PlayerRequestDto) model.Player
	GetPlayer(username string) model.Player
	GetAllPlayers() []model.Player
}

type playerService struct {
	playerRepository repository.PlayerRepository
}

func NewPlayerService(playerRepo repository.PlayerRepository) PlayerService {
	return &playerService{
		playerRepository: playerRepo,
	}
}

func (p *playerService) CreatePlayer(playerDto request.PlayerRequestDto) model.Player {
	player := model.Player{
		Name:     playerDto.Name,
		Username: playerDto.Username,
	}
	res := p.playerRepository.CreatePlayer(player)
	return res
}

func (p *playerService) GetPlayer(username string) model.Player {
	return p.playerRepository.GetPlayerByUsername(username)
}

func (p *playerService) GetAllPlayers() []model.Player {
	return p.playerRepository.GetAllPlayers()
}
