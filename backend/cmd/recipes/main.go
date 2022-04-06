package main

import (
	"log"

	"github.com/sflewis2970/recipes/middleware/corsconfig"
	"github.com/sflewis2970/recipes/routes"
)

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile)

	corsConfig := corsconfig.SetupCors()
	router := routes.SetupRouter(corsConfig)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}
