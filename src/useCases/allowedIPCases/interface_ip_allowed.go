/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package allowedIPCases

import (
	"github.com/gin-gonic/gin"	
)

/**  Interface */
type InterfaceDefaultUseCases interface {
	BodyGetData(c *gin.Context)
}

/**
	Method of interface
*/
func BodyGetData(p InterfaceDefaultUseCases , c *gin.Context) {
	p.BodyGetData(c)
}