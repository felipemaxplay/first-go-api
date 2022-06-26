package repository

import (
	"github.com/felipemaxplay/first-go-api/src/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlayerRepository interface {
	GetPlayerByUsername(username string) (model.Player, error)
	GetAllPlayers() ([]model.Player, error)
	CreatePlayer(player model.Player) (model.Player, error)
	UpdatePlayer(username string) (model.Player, error)
	DeletePlayer(username string) error
}

type playerConnection struct {
	connection *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) PlayerRepository {
	return &playerConnection{
		connection: db,
	}
}

func (p *playerConnection) GetPlayerByUsername(username string) (model.Player, error) {
	var player model.Player
	res := p.connection.Take(&player, "username = ?", username)
	if res.Error != nil {
		return player, res.Error
	}

	return player, nil
}

func (p *playerConnection) GetAllPlayers() ([]model.Player, error) {
	var players []model.Player
	p.connection.Find(&players)
	return players, nil
}

func (p *playerConnection) CreatePlayer(player model.Player) (model.Player, error) {
	player.ID = uuid.NewString()
	p.connection.Save(&player)
	return player, nil
}

func (p *playerConnection) UpdatePlayer(username string) (model.Player, error) {
	var player model.Player
	res := p.connection.Take(&player, "username = ?", username)
	if res != nil {
		return player, res.Error
	}

	p.connection.Save(&player)
	return player, nil
}

func (p *playerConnection) DeletePlayer(username string) error {
	var player model.Player
	res := p.connection.Take(&player, "username = ?", username)
	if res != nil {
		return res.Error
	}

	p.connection.Delete(&player)
	return nil
}
