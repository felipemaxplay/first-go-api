package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
