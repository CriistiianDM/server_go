/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package allowedIPCases

import (
	"server_go/src/domain/utils"
	"server_go/src/domain/allowedIPDomain"	

	"github.com/gin-gonic/gin"
)

var (
	instanceDom = allowedIPDomain.AllowedIPStruct{}
)

type UseCasesallowedIP struct{}

/**
   Return body of request
*/
func (p UseCasesallowedIP) BodyGetData(c *gin.Context) {
	requestParamns := utils.StateDefaultReq()
	allowedIPDomain.DomGetData(instanceDom,&requestParamns)
	utils.ResponseControlGeneral(c, requestParamns)
}