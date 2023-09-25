package main

import (
	"log"
	"net/http"
	"sensei/webservice-gin/controller"
	"sensei/webservice-gin/database"
	"sensei/webservice-gin/models"
	"sensei/webservice-gin/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	probing "github.com/prometheus-community/pro-bing"
)

func main() {

	loadEnv()
	loadDatabase()

	router := gin.Default()
	router.GET("/records", getPingStats)
	router.GET("/records/:id", getPingStatsByID)
	// router.POST("/record", postPingStats)
	router.POST("/ping", pingRunner)
	router.Run("localhost:8080")
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.PingRecord{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatal("Error Loading the Environment")
	}
}

func getPingStats(c *gin.Context) {
	records := controller.GetAllPingRecords()
	c.IndentedJSON(http.StatusOK, gin.H{"data": records})
}

// func postPingStats(stats models.PingRecord) (c *gin.Context) {
// 	createdRecord, err := controller.CreatePingRecord(stats)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}

// 	c.IndentedJSON(http.StatusCreated, gin.H{"success": createdRecord})
// }

func getPingStatsByID(c *gin.Context) {
	id := c.Param("id")
	stats, err := controller.GetRecordByID(id, c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No ping record exists with the given ID"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"success": stats})
}

func pingRunner(c *gin.Context) {
	resultChan := make(chan *probing.Statistics)
	errorChan := make(chan error)
	var request models.PingRequest
	c.BindJSON(&request)

	gameId, err := uuid.Parse(request.GameID)
	if err != nil {
		panic(err)
	}
	providerId, err := uuid.Parse(request.Provider)
	if err != nil {
		panic(err)
	}

	go controller.AsyncRunPing(&request, resultChan, errorChan)

	select {
	case result := <-resultChan:
		pingRecord := &models.PingRecord{
			MinLatency: utils.ConvertDurationToMs(result.MinRtt).String(),
			MaxLatency: utils.ConvertDurationToMs(result.MaxRtt).String(),
			AvgLatency: utils.ConvertDurationToMs(result.AvgRtt).String(),
			Ping:       utils.ConvertDurationToMs(result.StdDevRtt).String(),
			PacketLoss: result.PacketLoss,
			ServerIP:   result.Addr,
			GameID:     gameId,
			ProviderID: providerId,
		}

		pingRecord.Save()
		fullResponse := &models.PingResponse{
			Req:   &request,
			Probe: pingRecord,
		}

		if err != nil {
			panic(err.Error)
		}

		c.JSON(http.StatusOK, gin.H{"success": fullResponse})
	case err := <-errorChan:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
	}

	close(resultChan)
	close(errorChan)
}

//TODO Add tests for all methods
//TODO Document all endpoints. Try adding Swagger for this
//TODO Get IP, ISP and Game to test from the User. Run Pinger with given IP and fill in the result to the PingRecord Type
