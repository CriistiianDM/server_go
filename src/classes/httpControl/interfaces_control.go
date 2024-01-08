/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
*/
package httpControl

/**
  * Define the methods of the router
*/
type InterfaceControlHttp interface {
	GetParamsRequest() map[string]string
	GetBodyRequest() []byte
}

/**
  * Return all params in request
  * @return map[string]string
*/
func GetParamsRequest(p InterfaceControlHttp) map[string]string {
	return p.GetParamsRequest()
}

/**
  * Return Body of request
  * @return []byte
*/
func GetBodyRequest(p InterfaceControlHttp) []byte {
	return p.GetBodyRequest()
}