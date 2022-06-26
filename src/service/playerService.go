package service

import (
	"github.com/felipemaxplay/first-go-api/src/http/data/request"
	"github.com/felipemaxplay/first-go-api/src/model"
	"github.com/felipemaxplay/first-go-api/src/repository"
)

type PlayerService interface {
	CreatePlayer(playerDto request.PlayerRequestDto) (*model.Player, error)
	GetPlayer(username string) (*model.Player, error)
	GetAllPlayers() (*[]model.Player, error)
}

type playerService struct {
	playerRepository repository.PlayerRepository
}

func NewPlayerService(playerRepo repository.PlayerRepository) PlayerService {
	return &playerService{
		playerRepository: playerRepo,
	}
}

func (p *playerService) CreatePlayer(playerDto request.PlayerRequestDto) (*model.Player, error) {
	player := model.Player{
		Name:     playerDto.Name,
		Username: playerDto.Username,
	}
	res, err := p.playerRepository.CreatePlayer(player)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (p *playerService) GetPlayer(username string) (*model.Player, error) {
	player, err := p.playerRepository.GetPlayerByUsername(username)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (p *playerService) GetAllPlayers() (*[]model.Player, error) {
	players, err := p.playerRepository.GetAllPlayers()
	if err != nil {
		return nil, err
	}
	return &players, nil
}
