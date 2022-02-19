package corsconfig

import (
	"time"

	"github.com/gin-contrib/cors"
)

const PreFlightCacheLimit = 12

func SetupCors() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"Content-Type"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = PreFlightCacheLimit * time.Hour

	return corsConfig
}
