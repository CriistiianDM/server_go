/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
*/

package routesCompany

import (
	"server_go/repository/query"
	"database/sql"
)

var (
	sqlInstance = query.GeneralQuery{}
)

/**
  * Define the Company Routes
*/
type RoutesGeneral struct {}

/**
  * Return all routes of company
*/
func (r RoutesGeneral) GetCompanyRoutes() (*sql.Rows, error) {
	_query := "SELECT * FROM routes_company"
	return query.Query(sqlInstance, _query)
}