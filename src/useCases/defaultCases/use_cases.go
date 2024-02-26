/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package defaultCases

import (
	"server_go/src/domain/utils"

	"github.com/gin-gonic/gin"
)

var (
	requestParamns = utils.StateDefaultReq()
)

type UseCasesDefault struct{}

/**
   Return body of request
*/
func (p UseCasesDefault) BodyGetData(c *gin.Context) {
	utils.ResponseControlGeneral(c, requestParamns)
}