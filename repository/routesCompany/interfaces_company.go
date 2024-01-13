/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
*/

package routesCompany

import (
	_ "database/sql"
	_ "fmt"
)

/**
  * Define the methods of General Company Routes
*/
type InterfaceCompanyRoutes interface {
	GetCompanyRoutes() (map[string]interface{}, error)
}

/**
  * Return all routes of company
*/
func GetCompanyRoutes(p InterfaceCompanyRoutes) (map[string]interface{}, error) {
	return p.GetCompanyRoutes()
}