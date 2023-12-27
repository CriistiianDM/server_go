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
)

/**
   * Search in user in database
   * @param user
   */
func SearchOneUser(c *gin.Context) {
   fmt.Println("¡Buscando usuario!")
   fmt.Println(c)
   c.JSON(200, gin.H{"mensaje": "¡Perfil de usuario desde el controlador!"})
}