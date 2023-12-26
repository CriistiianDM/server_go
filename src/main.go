package main

// Librerary import
import (
	"github.com/gin-gonic/gin"
	"server_go/src/routes"
)

func main() {
	// Initialize the route
	router := gin.Default()

	// Create a group of routes
	apiGroup := router.Group("/api")
	routes.ConfigurarAPIRoutes(apiGroup)

	router.Run(":8080")
}