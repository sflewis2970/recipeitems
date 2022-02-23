package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sflewis2970/recipes/models"
)

const (
	NumberOfGroups = 1
	driverName     = "mysql"
)

var recipeMutex sync.Mutex

func GetRecipes(c *gin.Context) {
	// Call database driver to save record
	dbModel := models.New(driverName)
	recipeList, err := dbModel.GetRecipes()

	if err != nil {
		errMsg := err.Error()
		fmt.Println("database error: ", errMsg)

		// Send failed response to client
		recipeList := append(recipeList, models.Recipe{Message: errMsg})
		c.IndentedJSON(http.StatusInternalServerError, recipeList)
	} else {
		// Send success response to client
		c.IndentedJSON(http.StatusOK, recipeList)
	}
}

func CreateRecipe(c *gin.Context) {
	// Use mutex here so that when the data store is updated
	// only one resource can perform the write at a time
	recipeMutex.Lock()
	defer recipeMutex.Unlock()

	newRecipe := models.Recipe{}

	// extract request data from request
	extractRequestData(c, &newRecipe)

	// When a new recipe is created, a new UUID is generated
	// to serve as the unique ID
	uuid := uuid.New().String()

	if uuid != "" {
		uuid = buildUUID(uuid, "-", NumberOfGroups)
		newRecipe.Recipe_ID = uuid

		// Call database driver to save record
		dbModel := models.New(driverName)
		sqlRow := dbModel.AddRecipe(newRecipe)

		if sqlRow.Err() != nil {
			errMsg := sqlRow.Err().Error()
			fmt.Println("Error creating new recipe, with error", errMsg)
			newRecipe.Message = errMsg
			c.IndentedJSON(http.StatusInternalServerError, newRecipe)
		} else {
			c.IndentedJSON(http.StatusOK, newRecipe)
		}
	} else {
		newRecipe.Message = "Could not generate valid ID"
		c.IndentedJSON(http.StatusInternalServerError, newRecipe)
	}
}

func GetRecipe(c *gin.Context) {
	recipeID := getRecipeIDFromContext(c)

	if recipeID != "" {
		dbModel := models.New(driverName)
		recipe, err := dbModel.GetRecipe(recipeID)

		if err != nil {
			errMsg := err.Error()
			fmt.Println("Error getting recipe, with error", errMsg)
			recipe.Message = errMsg
			c.IndentedJSON(http.StatusInternalServerError, recipe)
		} else {
			c.IndentedJSON(http.StatusOK, recipe)
		}
	} else {
		recipe := models.Recipe{Message: "Request contains an invalid ID"}
		c.IndentedJSON(http.StatusBadRequest, recipe)
	}
}

func UpdateRecipe(c *gin.Context) {
	// Use mutex here so that when the data store is updated
	// only one resource can perform the write at a time
	recipeMutex.Lock()
	defer recipeMutex.Unlock()

	// extract request data from request
	var updatedRecipe models.Recipe
	extractRequestData(c, &updatedRecipe)

	// If the request contains a bad recipe ID,
	// return an invalid ID error message
	if updatedRecipe.Recipe_ID != "" {
		dbModel := models.New(driverName)
		err := dbModel.UpdateRecipe(updatedRecipe)

		if err != nil {
			errMsg := err.Error()
			fmt.Println("Error updating recipe, with error: ", errMsg)
			updatedRecipe.Message = errMsg
			c.IndentedJSON(http.StatusInternalServerError, updatedRecipe)

			return
		} else {
			recipe, _ := dbModel.GetRecipe(updatedRecipe.Recipe_ID)

			if recipe.Recipe_ID != updatedRecipe.Recipe_ID || recipe.Name != updatedRecipe.Name ||
				recipe.Ingredients != updatedRecipe.Ingredients || recipe.Instructions != updatedRecipe.Instructions ||
				recipe.Opened != updatedRecipe.Opened {
				updatedRecipe.Message = "Updated Failed"
				c.IndentedJSON(http.StatusInternalServerError, updatedRecipe)
			} else {
				c.IndentedJSON(http.StatusOK, updatedRecipe)
			}

			return
		}
	}

	updatedRecipe.Message = "Invalid ID"
	c.IndentedJSON(http.StatusBadRequest, updatedRecipe)
}

func OptionsRecipe(c *gin.Context) {
}

func DeleteRecipe(c *gin.Context) {
	// Use mutex here so that when the data store is updated
	// only one resource can perform the write at a time
	recipeMutex.Lock()
	defer recipeMutex.Unlock()

	recipeID := getRecipeIDFromContext(c)
	dbModel := models.New(driverName)
	deleteRecipe, err := dbModel.GetRecipe(recipeID)
	if err != nil {
		errMsg := err.Error()
		fmt.Println("Error finding recipe to delete, with error: ", errMsg)
		deleteRecipe.Message = errMsg
		c.IndentedJSON(http.StatusNotFound, deleteRecipe)
		return
	}

	// If the request contains a bad recipe ID,
	// return an invalid ID error message
	if recipeID != "" && recipeID == deleteRecipe.Recipe_ID {
		err = dbModel.DeleteRecipe(deleteRecipe.Recipe_ID)

		if err != nil {
			errMsg := err.Error()
			fmt.Println("Error deleting record, with error: ", errMsg)
			deleteRecipe.Message = errMsg
			c.IndentedJSON(http.StatusInternalServerError, deleteRecipe)
		} else {
			recipe, _ := dbModel.GetRecipe(deleteRecipe.Recipe_ID)

			if recipe.Recipe_ID != "" {
				deleteRecipe.Message = "Delete Failed"
				c.IndentedJSON(http.StatusInternalServerError, deleteRecipe)
			} else {
				c.IndentedJSON(http.StatusOK, deleteRecipe)
			}
		}

		return
	}

	deleteRecipe.Message = "Invalid ID"
	c.IndentedJSON(http.StatusBadRequest, deleteRecipe)
}

// Unexported functions
// Get Recipe ID from Context
func getRecipeIDFromContext(c *gin.Context) string {
	return c.Param("id")
}

// Extract Request Data
func extractRequestData(c *gin.Context, recipe *models.Recipe) error {
	err := c.BindJSON(recipe)

	if err != nil {
		return err
	}

	return nil
}

// Build UUID string
func buildUUID(uuid string, delimiter string, nbrOfGroups int) string {
	newUUID := ""

	uuidList := strings.Split(uuid, delimiter)
	for key, value := range uuidList {
		if key < nbrOfGroups {
			newUUID = newUUID + value
		}
	}

	return newUUID
}
