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
   "server_go/repository/productsQuerys"
)

var (
   newRequest = httpControl.HttpRequestControl{}
   instanceQuerys = productsQuerys.ProductsQuery{}
)

/**
   * Search in user in database
   * @param user
*/
func SearchOneUser(c *gin.Context) {
   newRequest.BodyRequest = c
   fmt.Println("newRequest", newRequest)
   params := httpControl.GetParamsRequest(newRequest)
   //sacar el id
   id := params["id"]
   fmt.Println("id", id)
   c.JSON(200, gin.H{"Congrulations": "New Year 2024, Happy New Year!"})
}

/**
   * Insert in product in database
*/
func Insertproduct(c *gin.Context) {
   newRequest.BodyRequest = c
   params := httpControl.GetBodyRequest(newRequest)
   _, err := productsQuerys.InserProductCompanys(instanceQuerys,params)
   status := 200
   _res :=  gin.H{"response": true , "message": "Product inserted"}
   
   if err != nil {
      status = 500
      _res = gin.H{"response": false , "message": "Error in insert product"}
   }

   c.JSON(status, _res)
}

/**
   * Not found
*/
func NotFound(c *gin.Context) {
   c.JSON(404, gin.H{"error": "No found : - 404 -:"})
}