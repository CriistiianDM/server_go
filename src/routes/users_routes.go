// rutas/api_routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"server_go/src/controls"
)

func ConfigurarAPIRoutes(router *gin.RouterGroup) {
	router.GET("/saludo", controls.SearchOneUser)
}
