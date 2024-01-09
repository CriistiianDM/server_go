/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2023
   * @license: GPL-3.0
*/
package main

// Librerary import
import (
	"server_go/repository/routesCompany"
	 _ "server_go/repository/query"
	"github.com/gin-gonic/gin"
	"server_go/src/routes"
	"server_go/db"
	"fmt"
	"sync"
)

/**
  * Main function
*/
func main() {
	_executeRun();
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

/**
  * Execute all goroutines
*/
func _executeRun() {
	// Esto es magia 
	// Literal tengo el control del tiempo de la concurrencia
	var wg sync.WaitGroup
		wg.Add(1)
		// Cuando acabe la rutina, sigen las demas instrucciones
		go db.Connect(&wg);
	wg.Wait()
	_initRoutesCompany();
}

func _initRoutesCompany() {
	fmt.Println("Init routes company")
	companyInstance := routesCompany.RoutesGeneral{}
	result, err := routesCompany.GetCompanyRoutes(companyInstance)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Result", result)
}