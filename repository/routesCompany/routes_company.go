/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
*/

package routesCompany

import (
	"server_go/repository/query"
	_ "database/sql"
	"fmt"
)

var (
	sqlInstance = query.GeneralQuery{}
	_query = "SELECT * FROM routes_company"
)

/**
  * Define the Company Routes
*/
type RoutesGeneral struct {}

type RoutesCompany struct {
	id int
	route string
	created_at string
	updated_at string
}

/**
  * Return all routes of company
*/
func (r RoutesGeneral) GetCompanyRoutes() (map[string]interface{}, error) {
	result, err := query.Query(sqlInstance, _query)
	mapData := make(map[string]interface{})

	if err == nil {
		defer result.Close()
		for result.Next() {
			instance := RoutesCompany{}
			err := result.Scan(&instance.id, &instance.route, &instance.created_at, &instance.updated_at)
			if err != nil {
				fmt.Println("Error", err)
			}
			mapData[instance.route] = instance
			fmt.Println("Instance", instance)
		}
	}

	return mapData, err
}