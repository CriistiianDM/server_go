/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
*/

package routesCompany

import (
	"database/sql"
	_ "fmt"
)

/**
  * Define the methods of General Company Routes
*/
type InterfaceCompanyRoutes interface {
	GetCompanyRoutes() (*sql.Rows, error)
}

/**
  * Return all routes of company
*/
func GetCompanyRoutes(p InterfaceCompanyRoutes) (*sql.Rows, error) {
	return p.GetCompanyRoutes()
}