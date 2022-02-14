package main

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type recipe struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Ingredients  string `json:"ingredients"`
	Instructions string `json:"instructions"`
}

const (
	PreFlightCacheLimit = 12
	NumberOfGroups      = 3
)

var recipes = []recipe{}
var recipeMutex sync.Mutex

// Exported functions
func GetRecipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, recipes)
}

func CreateRecipe(c *gin.Context) {
	// Use mutex here so that when the data store is updated
	// only one resource can perform the write at a time
	recipeMutex.Lock()
	defer recipeMutex.Unlock()

	var newRecipe recipe

	// extract request data from request
	extractRequestData(c, &newRecipe)

	// When a new recipe is created, a new UUID is generated
	// to serve as the unique ID
	uuidStr := uuid.New().String()

	if uuidStr != "" {
		uuidStr = parseUUIDStr(uuidStr, "-", NumberOfGroups)
		newRecipe.ID = uuidStr
	}

	// Update data store. For now use a local data store
	// Evenually a MySQL database will be used
	recipes = append(recipes, newRecipe)

	c.IndentedJSON(http.StatusOK, newRecipe)
}

func GetRecipe(c *gin.Context) {
	recipe := lookupRecipeWithContext(c)

	if recipe.ID != "" {
		c.IndentedJSON(http.StatusOK, recipe)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipe not found"})
}

func UpdateRecipe(c *gin.Context) {
	// Use mutex here so that when the data store is updated
	// only one resource can perform the write at a time
	recipeMutex.Lock()
	defer recipeMutex.Unlock()

	var updatedRecipe recipe

	// extract request data from request
	extractRequestData(c, &updatedRecipe)

	if updatedRecipe.ID != "" {
		// Find recipe in data store
		recipe := lookupRecipeByID(updatedRecipe.ID)

		if recipe.ID != "" {
			recipe = updatedRecipe

			c.IndentedJSON(http.StatusOK, recipe)
			return
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipe not found"})
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

	recipe := lookupRecipeWithContext(c)

	if recipe.ID != "" {
		pos := lookupRecipePosByID(recipe.ID)

		if pos != -1 {
			recipes := append(recipes[0:pos], recipes[pos+1:]...)
			c.IndentedJSON(http.StatusOK, recipes)
			return
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipe not found"})
}

// Unexported functions
// Get recipe using parameter ID in Context parameter
func lookupRecipeWithContext(c *gin.Context) recipe {
	id := c.Param("id")

	for _, recipe := range recipes {
		if recipe.ID == id {
			return recipe
		}
	}

	return recipe{}
}

// Get recipe by ID
func lookupRecipeByID(id string) recipe {
	for _, recipe := range recipes {
		if recipe.ID == id {
			return recipe
		}
	}

	return recipe{}
}

// Get recipe positon by ID
func lookupRecipePosByID(id string) int {
	for idx, recipe := range recipes {
		if recipe.ID == id {
			return idx
		}
	}

	return -1
}

// Extract Request Data
func extractRequestData(c *gin.Context, recipe *recipe) error {
	err := c.BindJSON(recipe)

	if err != nil {
		return err
	}

	return nil
}

// Parse UUID string
func parseUUIDStr(uuidStr string, delimiter string, nbrOfGroups int) string {
	newUUIDStr := ""

	uuidList := strings.Split(uuidStr, delimiter)
	for idx, uuid := range uuidList {
		if idx < nbrOfGroups {
			newUUIDStr = newUUIDStr + uuid
		}
	}

	return newUUIDStr
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
