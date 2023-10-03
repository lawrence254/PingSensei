package models

import (
	"sensei/webservice-gin/database"
	"time"

	"github.com/google/uuid"
)

type Providers struct {
	ID        uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	IspName   string       `gorm:"type:string;not null" json:"provider_name"`
	Country   string       `gorm:"type:string;not null" json:"counrty"`
	Records   []PingRecord `gorm:"foreignKey:ProviderID" json:"ping_stats"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ProviderRequest struct {
	IspName string `json:"provider_name"`
	Country string `json:"country"`
}

func (provider *Providers) Save() (*Providers, error) {
	err := database.Database.Create(&provider).Error
	if err != nil {
		return &Providers{}, err
	}

	return provider, nil
}
