package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sflewis2970/recipes/api/services/database"
	"github.com/sflewis2970/recipes/api/services/recipes"
)

const (
	PreFlightCacheLimit = 12
	NumberOfGroups      = 3
)

// var recipeList []recipes.Recipe
var recipeMutex sync.Mutex

// Exported functions
func GetRecipes(c *gin.Context) {
	// Call database driver to save record
	recipeList, err := database.GetRecipes()

	if err != nil {
		fmt.Println("database error: ", err.Error())

		// Send failed response to client
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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

	newRecipe := recipes.Recipe{}

	// extract request data from request
	extractRequestData(c, &newRecipe)

	// When a new recipe is created, a new UUID is generated
	// to serve as the unique ID
	uuid := uuid.New().String()

	if uuid != "" {
		uuid = buildUUID(uuid, "-", NumberOfGroups)
		newRecipe.Recipe_ID = uuid

		// Call database driver to save record
		sqlRow := database.AddRecipe(newRecipe)

		if sqlRow.Err() != nil {
			fmt.Println("Error creating new recipe, with error", sqlRow.Err().Error())
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": sqlRow.Err().Error()})
		} else {
			c.IndentedJSON(http.StatusOK, newRecipe)
		}
	} else {
		c.IndentedJSON(http.StatusInternalServerError, recipes.Recipe{})
	}
}

func GetRecipe(c *gin.Context) {
	recipeID := getRecipeIDFromContext(c)

	if recipeID != "" {
		recipe, err := database.GetRecipe(recipeID)

		if err != nil {
			fmt.Println("Error getting recipe, with error", err.Error())
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, recipe)
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
	}
}

func UpdateRecipe(c *gin.Context) {
	// Use mutex here so that when the data store is updated
	// only one resource can perform the write at a time
	recipeMutex.Lock()
	defer recipeMutex.Unlock()

	// extract request data from request
	var updatedRecipe recipes.Recipe
	extractRequestData(c, &updatedRecipe)

	// If the request contains a bad recipe ID,
	// return an invalid ID error message
	if updatedRecipe.Recipe_ID != "" {
		err := database.UpdateRecipe(updatedRecipe)

		if err != nil {
			fmt.Println("Error updating recipe, with error: ", err.Error())
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			recipe, _ := database.GetRecipe(updatedRecipe.Recipe_ID)

			if recipe.Recipe_ID != updatedRecipe.Recipe_ID || recipe.Name != updatedRecipe.Name &&
				recipe.Ingredients != updatedRecipe.Ingredients || recipe.Instructions != updatedRecipe.Instructions {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Updated Failed"})
				return
			}

			c.IndentedJSON(http.StatusOK, gin.H{"message": "Update Successful"})
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
	}
}

func OptionsRecipe(c *gin.Context) {
	// c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	// c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	// c.Header("Access-Control-Allow-Headers", "Content-Type")
}

func DeleteRecipe(c *gin.Context) {
	// Use mutex here so that when the data store is updated
	// only one resource can perform the write at a time
	recipeMutex.Lock()
	defer recipeMutex.Unlock()

	recipeID := getRecipeIDFromContext(c)

	// If the request contains a bad recipe ID,
	// return an invalid ID error message
	if recipeID != "" {
		err := database.DeleteRecipe(recipeID)

		if err != nil {
			fmt.Println("Error deleting record, with error: ", err.Error())
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			recipe, _ := database.GetRecipe(recipeID)

			if recipe.Recipe_ID != "" {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Delete Failed"})
				return
			}

			c.IndentedJSON(http.StatusOK, gin.H{"message": "Delete Successful"})
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
	}
}

// Unexported functions
// Get Recipe ID from Context
func getRecipeIDFromContext(c *gin.Context) string {
	return c.Param("id")
}

// Extract Request Data
func extractRequestData(c *gin.Context, recipe *recipes.Recipe) error {
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

func main() {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"Content-Type"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = PreFlightCacheLimit * time.Hour

	router.Use(cors.New(corsConfig))

	router.GET("/recipes", GetRecipes)
	router.POST("/recipe", CreateRecipe)
	router.GET("/recipe/:id", GetRecipe)
	router.PUT("/recipe", UpdateRecipe)
	router.OPTIONS("/recipe", OptionsRecipe)
	router.DELETE("/recipe/:id", DeleteRecipe)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}
