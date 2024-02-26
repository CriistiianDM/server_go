/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
 */

package defaultControl

// Librerary import
import (
	"server_go/src/useCases/defaultCases"

	"github.com/gin-gonic/gin"
)

// Vars
var (
	instanceDefault = defaultCases.UseCasesDefault{}
)

/**
 * Return Follower per user
 */
func GetData(c *gin.Context) {
	defaultCases.BodyGetData(instanceDefault, c)
}