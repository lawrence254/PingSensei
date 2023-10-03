package main

import (
	"log"
	"net/http"
	"sensei/webservice-gin/database"
	"sensei/webservice-gin/management"
	"sensei/webservice-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()

	router := gin.Default()
	router.POST("/management/games", addGame)
	router.PATCH("/management/games/:id", updateGame)
	router.GET("/management/games", getAllGames)
	router.GET("/management/games/:id", getGameById)

	router.Run("localhost:8080")
}

// loadDatabase() is used to initialize the database connection and setup the automigration of the specified models using gorm.
func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.Games{}, &models.Providers{})
}

// loadEnv() is used to initialize the environment to be used. Options are local,development or production
func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error Loading Environment.")
	}
}

// addGame() is a function that is used to add a given game to the database table. It uses the gin.Context and this must be provided.
func addGame(c *gin.Context) {

	result, err := management.CreateGame(c)

	if err != nil {
		log.Fatal("Failed to create game record. The error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusAccepted, gin.H{"success": result})

}

// updateGame() will update the given game based on the game id provided.
func updateGame(c *gin.Context) {
	result, err := management.UpdateGameDetails(c)

	if err != nil {
		log.Fatal("Failed to update game record. The error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusAccepted, gin.H{"success": result})

}

// getAllGames() will return a list of all the Games that have been saved to the database.
func getAllGames(c *gin.Context) {
	result := management.GetAllGames()

	if len(result) < 1 {
		c.JSON(http.StatusOK, gin.H{"Failed":"There are no saved games. Please create one to get the details"})
		return
	}else{
		c.JSON(http.StatusOK, gin.H{"success": result})
	}

}

// getGameById() will return a single game that has an ID matching the id provided.
func getGameById(c *gin.Context) {

	result, err := management.GetGameByID(c)

	if err != nil {
		log.Fatal("Failed to retrive the game record. The error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusAccepted, gin.H{"success": result})
}

// addProvider() will create and return a single provider.
func addProvider(c *gin.Context) {
	result, err := management.CreateProvider(c)

	if err != nil {
		log.Fatal("Failed to save the given provider. The error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"success": result})
}

// removeProvider() will soft delete a record based on the ID and hide the corresponding provider from future searches.
func removeProvider(c *gin.Context) {
	result := management.DeleteProvider(c)

	if result != nil {
		log.Fatal("Failed to delete provider. The error is", result.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Failed": result.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"success": "Successfully deleted the provider"})
}

// updateProvider() will update the provider based on the id that has been submitted
func updateProvider(c *gin.Context) {
	result, err := management.UpdateProvider(c)

	if err != nil {
		log.Fatal("Failed to update the provider with the given details. The error details are :", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"success": result})
}

// getAllProviders() will return all the providers currently save in the database.
func getAllProviders(c *gin.Context) {
	result := management.GetAllProviders()

	if len(result) < 1 {
		c.JSON(http.StatusOK, gin.H{"success": "No records found"})
	}

	c.JSON(http.StatusOK, gin.H{"success": result})
}

// getProviderById() will return the requested provider data if available
func getProviderById(c *gin.Context) {
	result, err := management.GetProviderByID(c)

	if err != nil {
		log.Fatal("Failed to fetch the requested provider. The error is: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"success": result})

}
