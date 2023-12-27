/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2023
   * @license: GPL-3.0
*/
package httpRequest

// Librerary import
import _ "github.com/gin-gonic/gin"

/* Interface declaration */

type HttpControls interface {
  GetBody() string
}
