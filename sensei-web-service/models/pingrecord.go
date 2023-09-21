package models

import (
	"sensei/webservice-gin/database"
	"time"

	"github.com/google/uuid"
)

type PingRecord struct {
	ID         uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	MinLatency string        `gorm:"size:50;not null" json:"minlatency"`
	MaxLatency string        `gorm:"size:50;not null" json:"maxlatency"`
	AvgLatency string        `gorm:"size:50;not null" json:"avglatency"`
	Ping       time.Duration `gorm:"size:50;not null" json:"ping"`
	PacketLoss float64       `gorm:"size:50;not null" json:"packetloss"`
	GameID     uuid.UUID     `gorm:"not null;type:uuid;default:uuid_generate_v4()" json:"gameid"`
	ProviderID uuid.UUID     `gorm:"not null;type:uuid;default:uuid_generate_v4()" json:"providerid"`
	ServerIP   string        `gorm:"not null" json:"serverip"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func (record *PingRecord) Save() (*PingRecord, error) {
	err := database.Database.Create(&record).Error

	if err != nil {
		return &PingRecord{}, err
	}
	return record, nil
}
