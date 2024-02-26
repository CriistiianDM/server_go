/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package defaultCases

import (
	"server_go/src/domain/utils"
	"server_go/src/domain/defaultDomain"	

	"github.com/gin-gonic/gin"
)

var (
	instanceDom = defaultDomain.DefaultDomain{}
)

type UseCasesDefault struct{}

/**
   Return body of request
*/
func (p UseCasesDefault) BodyGetData(c *gin.Context) {
	requestParamns := utils.StateDefaultReq()
	defaultDomain.DomGetData(instanceDom,&requestParamns)
	utils.ResponseControlGeneral(c, requestParamns)
}