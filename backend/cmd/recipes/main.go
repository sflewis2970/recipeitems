package main

import (
	"github.com/sflewis2970/recipes/middleware/corsconfig"
	"github.com/sflewis2970/recipes/routes"
)

func main() {
	corsConfig := corsconfig.SetupCors()
	router := routes.SetupRouter(corsConfig)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}
