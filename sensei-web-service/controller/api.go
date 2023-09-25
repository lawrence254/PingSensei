package controller

import (
	"sensei/webservice-gin/database"
	"sensei/webservice-gin/models"
	"sensei/webservice-gin/utils"

	"github.com/gin-gonic/gin"
	probing "github.com/prometheus-community/pro-bing"
)

func CreatePingRecord(stats models.PingRecord) (*models.PingRecord, error) {
	var newPingData models.PingRecord

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

func GetAllPingRecords() ([]models.PingRecord){
	var records []models.PingRecord
	database.Database.Find(&records)
	return records
}

func AsyncRunPing(req *models.PingRequest, resultChan chan<- *probing.Statistics, errorChan chan <- error){
	
	go func() {
		result, err := utils.PingServer(req.ServerIP)
		if err != nil {
			errorChan <- err
		}else{
		resultChan <- result
		}
	}()
}