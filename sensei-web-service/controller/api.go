package controller

import (
	"sensei/webservice-gin/database"
	"sensei/webservice-gin/models"
	"sensei/webservice-gin/utils"

	"github.com/gin-gonic/gin"
	probing "github.com/prometheus-community/pro-bing"
)

// CreatePingRecord creates a ping record in the database for the given IP adderss contained within the PingRecord type.
func CreatePingRecord(stats models.PingRecord) (*models.PingRecord, error) {
	var newPingData models.PingRecord

	record, err := newPingData.Save()
	if err != nil {
		return &newPingData, err
	}
	return record, nil
}

// GetRecordByID returns the ping record corresponding to the given id string which is a UUID.
func GetRecordByID(id string, c *gin.Context) (*models.PingRecord, error) {
	var record models.PingRecord
	err := database.Database.First(&record, "id = ?", id).Error

	if err != nil {
		return &record, err
	}

	return &record, nil
}

// GetAllPingRecords returns all the ping records that have been captured.
func GetAllPingRecords() []models.PingRecord {
	var records []models.PingRecord
	database.Database.Find(&records)
	return records
}

// AsyncRunPing is an asynchronous function that will attempt to run a ping command for the given server address using the [github.com/prometheus-community/pro-bing] package.
func AsyncRunPing(req *models.PingRequest, resultChan chan<- *probing.Statistics, errorChan chan<- error) {

	go func() {
		result, err := utils.PingServer(req.ServerIP)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- result
		}
	}()
}
