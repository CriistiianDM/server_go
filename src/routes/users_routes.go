/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2023
   * @license: GPL-3.0
*/
package routes

// Librerary import
import (
	  "server_go/src/controls"
     "server_go/src/interfaces/httpRequest"
)

/* Router type declaration */
func InitializeApiRoutes(router httpRequest.HttpRouter , allRoutes map[string]interface{}) {
   if allRoutes != nil || len(allRoutes) != 0 {
      for key := range allRoutes {
         router.GET(key, httpRequest.HandleSync(controls.SearchOneUser))
      }
   }
}
