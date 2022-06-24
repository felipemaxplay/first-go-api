package model

type Player struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(40); not null; default:null"`
	Username string `json:"username" gorm:"type:varchar(60); not null; unique; default:null"`
}
