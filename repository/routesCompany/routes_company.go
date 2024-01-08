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
	"fmt"
)

/**
  * Define the Company Routes
*/
type RoutesGeneral struct {
	Query query.InterfaceQuery
}

/**
  * Return all routes of company
*/
func (r *RoutesGeneral) GetCompanyRoutes() (*sql.Rows, error) {
	QUERY := "SELECT * FROM routes_company"
	result , err := r.Query.Query(QUERY)
	fmt.Println(result)
	fmt.Println("Queryyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy ***********************************")
	return result, err
}