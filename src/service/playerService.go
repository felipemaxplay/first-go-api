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
	UpdatePlayer(username string, playerDto request.PlayerRequestDto) (*model.Player, error)
	DeletePlayer(username string) error
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

func (p *playerService) UpdatePlayer(username string, playerDto request.PlayerRequestDto) (*model.Player, error) {
	player, err := p.playerRepository.GetPlayerByUsername(username)
	if err != nil {
		return nil, err
	}

	player.Name = playerDto.Name
	player.Username = playerDto.Username

	player, err = p.playerRepository.UpdatePlayer(player)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (p *playerService) DeletePlayer(username string) error {
	err := p.playerRepository.DeletePlayer(username)
	if err != nil {
		return err
	}
	return nil
}
