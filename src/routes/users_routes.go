/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2023
   * @license: GPL-3.0
*/
package routes

// Librerary import
import (
	_ "github.com/gin-gonic/gin"
	  "server_go/src/controls"
     "server_go/src/interfaces/httpRequest"
)

/* Router type declaration */
func ConfigurarAPIRoutes(router httpRequest.HttpRouter) {
   router.GET("/new-year/:name/:id", controls.SearchOneUser)
}
