package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PingRecord struct {
	ID         string        `json:"id"`
	MinLatency string        `json:"minlatency"`
	MaxLatency string        `json:"maxlatency"`
	AvgLatency string        `json:"avglatency"`
	Ping       time.Duration `json:"ping"`
	PacketLoss float64       `json:"packetloss"`
	GameID     string        `json:"gameid"`
	ProviderID string        `json:"providerid"`
	ServerIP   string        `json:"serverip"`
	Date       string        `json:"date"`
}

var records = []PingRecord{
	{
		ID:         "1f767283-b3c4-4528-b658-09ac104a094c",
		MinLatency: "1990.341",
		MaxLatency: "2388.485",
		AvgLatency: "869.206",
		Ping:       20,
		PacketLoss: 1.41,
		GameID:     "b7efa95c-2e88-4ecb-a869-01606a32c2ce",
		ProviderID: "05c4ee56-dfc9-4cfc-bf5a-02371282ff0a",
		ServerIP:   "190.110.171.222",
		Date:       "2006-01-02",
	}}

func main() {
	router := gin.Default()
	router.GET("/records", getPingStats)
	router.GET("/records/:id", getPingStatsByID)
	router.POST("/record", postPingStats)
	router.Run("localhost:8080")
}
func getPingStats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, records)
}
func postPingStats(c *gin.Context) {
	var newPingData PingRecord

	if err := c.BindJSON(&newPingData); err != nil {
		return
	}

	records = append(records, newPingData)
	c.IndentedJSON(http.StatusCreated, newPingData)
}

func getPingStatsByID(c *gin.Context){
	id :=c.Param("id")
	for _, a := range records{
		if a.ID== id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}	
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"No ping records exist with the given ID"})
}