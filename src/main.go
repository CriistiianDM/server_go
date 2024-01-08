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
	"server_go/repository/query"
	"github.com/gin-gonic/gin"
	"server_go/src/routes"
	"server_go/db"
	"time"
	"fmt"
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
	go db.Connect()
	// classRoutes := routesCompany.RoutesGeneral{
	// 	Query: &query.InterfaceQuery{},
	// }
	// result, err := classRoutes.GetCompanyRoutes()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(result)
	time.Sleep(1 * time.Second)
}

func imprimirNumeros() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}