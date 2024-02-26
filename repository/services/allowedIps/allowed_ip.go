/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package allowedIps

// Librerary import
import (
	"server_go/repository/config"
	"server_go/repository/services/utils"
)

var (
	_query_allowed_ip = config.ConfigMapSql()["get_allowed_ips"]
)

/* Class DataUserEmail */
type AllowedIpsUtils struct {}

/**
  * Return all allowed ips
*/
func (p AllowedIpsUtils) GetAllowedIps(args ...interface{}) ([]map[string]interface{}, error) {
	return utils.ExecuteFileSql(_query_allowed_ip, args...);
}