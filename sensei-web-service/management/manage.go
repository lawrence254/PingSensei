package management

import (
	"sensei/webservice-gin/database"
	"sensei/webservice-gin/models"

	"github.com/gin-gonic/gin"
)

func CreateGame(c *gin.Context) (*models.Games, error) {
	var game models.Games

	if err := c.ShouldBindJSON(&game); err != nil {
		return &game, err
	}
	result, err := game.Save()

	if err != nil {
		return &game, err
	}
	return result, nil
}

func GetAllGames() []models.Games {
	var games []models.Games

	database.Database.Find(&games)

	return games
}

func GetGameByID(c *gin.Context) (*models.Games, error) {
	var game models.Games
	id := c.Param("id")
	err := database.Database.Find(&game, "id = ?", id).Error

	if err != nil {
		return &game, err
	}

	return &game, nil
}

func UpdateGameDetails(c *gin.Context) (*models.Games, error) {

	var game models.Games

	id := c.Param("id")
	err := database.Database.Find(&game, "id = ?", id).Error

	if err != nil {
		return &game, err
	}

	var updateGame models.Games
	if bindErr := c.ShouldBindJSON(&updateGame); bindErr != nil {
		return &updateGame, bindErr
	}

	result, updateErr := updateGame.Save()

	if updateErr != nil {
		return &updateGame, updateErr
	}

	return result, nil
}

func CreateProvider(c *gin.Context) (*models.Providers, error) {
	var provider models.Providers

	if err := c.ShouldBindJSON(&provider); err != nil {
		return &provider, err
	}

	result, err := provider.Save()

	if err != nil {
		return &provider, err
	}

	return result, nil
}

func GetAllProviders() []models.Providers {
	var providers []models.Providers

	database.Database.Find(&providers)

	return providers
}

func GetProviderByID(c *gin.Context) (*models.Providers, error) {
	var provider models.Providers

	id := c.Param("id")

	err := database.Database.Find(&provider, "id = ?", id).Error

	if err != nil {
		return &provider, err
	}

	return &provider, nil
}

func UpdateProvider(c *gin.Context) (*models.Providers, error) {
	var existing models.Providers

	id := c.Param("id")
	err := database.Database.Find(&existing, "id = ?", id).Error
	if err != nil {
		return &existing, err
	}

	var updatedProvider models.Providers
	if bindErr := c.ShouldBindJSON(&updatedProvider); bindErr != nil {
		return &updatedProvider, bindErr
	}

	result, UpdateError := updatedProvider.Save()
	if UpdateError != nil {
		return &updatedProvider, UpdateError
	}

	return result, nil
}

func DeleteProvider(c *gin.Context) error{
	id := c.Param("id")

	if err := database.Database.Delete(&models.Providers{},id).Error; err != nil{
		return err
	}

	return nil
}