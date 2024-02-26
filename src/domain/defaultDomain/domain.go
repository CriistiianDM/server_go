/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package defaultDomain

import (
	"server_go/repository/services/allowedIps"
	"fmt"
)

var (
	instanceIp = allowedIps.AllowedIpsUtils{}
)

type DefaultDomain struct {}

/**
  * Return all allowed ips
*/
func (p DefaultDomain) DomGetData(args *map[string]interface{}) {
	if len(*args) > 0 {
		_response, is_error := allowedIps.GetAllowedIps(instanceIp, []interface{}{true}...)
		
		if is_error == nil {
			(*args)["data"] = _response;
			(*args)["status"] = 200;
			(*args)["message_default"] = "Success";
			(*args)["statusReq"] = true
		} else {
			fmt.Println("Error: " , is_error)
			(*args)["message_default"] = "Error get allowed ips";
		}

	} 
}