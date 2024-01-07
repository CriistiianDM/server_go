/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2023
   * @license: GPL-3.0
*/
package main

// Librerary import
import (
	"github.com/gin-gonic/gin"
	"server_go/src/routes"
)

/**
  * Main function
*/
func main() {
	_initServer();
}

/**
  * Initialize the server
*/
func _initServer() {
	// Initialize the route
	router := gin.Default()

	// Create a group of routes
	apiGroup := router.Group("/api")
	routes.InitializeApiRoutes(apiGroup)

	router.Run(":8080")
}