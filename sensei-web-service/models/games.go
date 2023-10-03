package models

import (
	"sensei/webservice-gin/database"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Games struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string         `gorm:"type:string" json:"game_name"`
	Publisher   string         `gorm:"type:string" json:"publisher_name"`
	IpAddresses pq.StringArray `gorm:"type:text[]" json:"ip_addresses"`
	Records 	[]PingRecord	`gorm:"foreignKey:GameID" json:"ping_stats"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func (game *Games) Save() (*Games, error){
	err := database.Database.Create(&game).Error

	if err != nil {
		return &Games{}, err
	}

	return game, nil
}