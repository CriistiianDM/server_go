/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2023
   * @license: GPL-3.0
*/
package controls;

import (
   "github.com/gin-gonic/gin"
    _ "fmt"
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
func GetProductsCompany(c *gin.Context) {
   newRequest.BodyRequest = c
   params := httpControl.GetBodyRequest(newRequest)
   result, err := instanceQuerys.GetProductsPerCompany(params)
   status := 500
   _res := gin.H{"response": false , "message": "Error get products"}

   if err == nil {
      status = 200
      _res = result
   }
   c.JSON(status, _res)
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