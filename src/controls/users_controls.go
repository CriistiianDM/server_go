/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2023
   * @license: GPL-3.0
*/
package controls;

import (
   "github.com/gin-gonic/gin"
   "fmt"
   "server_go/src/classes/httpControl"
)

/**
   * Search in user in database
   * @param user
   */
func SearchOneUser(c *gin.Context) {
   fmt.Println("¡Buscando usuario!")
   fmt.Println(*c)
   newRequest := httpControl.HttpRequestControl{
      BodyRequest: c,
   }
   fmt.Println(newRequest.GetParamsRequest())
   c.JSON(200, gin.H{"Congrulations": "New Year 2024, Happy New Year!"})
}