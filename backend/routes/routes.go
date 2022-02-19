package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sflewis2970/recipes/controllers"
)

func SetupRouter(corsConfig cors.Config) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(corsConfig))
	router.GET("/api/recipes", controllers.GetRecipes)
	router.POST("/api/recipes", controllers.CreateRecipe)
	router.GET("/api/recipes/:id", controllers.GetRecipe)
	router.PUT("/api/recipes", controllers.UpdateRecipe)
	router.OPTIONS("/api/recipes", controllers.OptionsRecipe)
	router.DELETE("/api/recipes/:id", controllers.DeleteRecipe)

	return router
}
