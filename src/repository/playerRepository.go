package repository

import (
	"github.com/felipemaxplay/first-go-api/src/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlayerRepository interface {
	GetPlayerByUsername(username string) model.Player
	GetAllPlayers() []model.Player
	CreatePlayer(player model.Player) model.Player
	UpdatePlayer(username string) model.Player
	DeletePlayer(username string)
}

type playerConnection struct {
	connection *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) PlayerRepository {
	return &playerConnection{
		connection: db,
	}
}

func (p *playerConnection) GetPlayerByUsername(username string) model.Player {
	var player model.Player
	p.connection.Find(&player, "username = ?", username)
	return player
}

func (p *playerConnection) GetAllPlayers() []model.Player {
	var players []model.Player
	p.connection.Find(&players)
	return players
}

func (p *playerConnection) CreatePlayer(player model.Player) model.Player {
	player.ID = uuid.NewString()
	p.connection.Save(&player)
	return player
}

func (p *playerConnection) UpdatePlayer(username string) model.Player {
	var player model.Player
	p.connection.Find(&player, "username = ?", username)
	p.connection.Save(&player)
	return player
}

func (p *playerConnection) DeletePlayer(username string) {
	var player model.Player
	p.connection.Find(&player, "username = ?", username)
	p.connection.Delete(&player)
}
