package controller

import (
	"sensei/webservice-gin/database"
	"sensei/webservice-gin/models"

	"github.com/gin-gonic/gin"
)

func CreatePingRecord(c *gin.Context) (*models.PingRecord, error) {
	var newPingData models.PingRecord

	if err := c.BindJSON(&newPingData); err != nil {
		return &newPingData, err
	}

	record, err := newPingData.Save()
	if err != nil {
		return &newPingData, err
	}
	return record, nil
}

func GetRecordByID(id string, c *gin.Context) (*models.PingRecord, error) {
	var record models.PingRecord
	err := database.Database.First(&record, "id = ?", id).Error

	if err != nil {
		return &record, err
	}

	return &record, nil
}

func GetAllPingRecords () ([]models.PingRecord){
	var records []models.PingRecord
	database.Database.Find(&records)
	return records
}
