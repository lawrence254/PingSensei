package main

import (
	"log"
	"net/http"
	"sensei/webservice-gin/controller"
	"sensei/webservice-gin/database"
	"sensei/webservice-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	loadEnv()
	loadDatabase()

	router := gin.Default()
	router.GET("/records", getPingStats)
	router.GET("/records/:id", getPingStatsByID)
	router.POST("/record", postPingStats)
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

func postPingStats(c *gin.Context) {
	createdRecord, err := controller.CreatePingRecord(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"success": createdRecord})
}

func getPingStatsByID(c *gin.Context) {
	id := c.Param("id")
	stats,err := controller.GetRecordByID(id, c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No ping record exists with the given ID"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"success": stats})
}

//TODO Add DELETE and UPDATE Methods for the records
//TODO Add tests for all methods
//TODO Document all endpoints. Try adding Swagger for this
