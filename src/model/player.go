package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	ID       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
