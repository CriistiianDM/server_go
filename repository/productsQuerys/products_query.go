/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
*/
package productsQuerys

import (
	"server_go/repository/query"
	"encoding/json"
	"fmt"
)

const (
	_query = "CALL insert_new_product_per_company($1,$2)"
)

type ProductsQuery struct {}

/**
  * Insert new product per company
*/
func (p ProductsQuery) InserProductCompanys(params map[string]interface{}) (bool, error) {
	// var 
	company, ok_company := params["company"]
	json_, ok_json := p.mapToJSON(params["json"])
	response := false
	var is_err error

	if ok_company && ok_json {
		result, err := query.Exec(query.GeneralQuery{},_query, company, json_)
		is_err = err
		if err == nil {
			fmt.Println("Result", result)
			response = !response
		}
	}

	return response, is_err
}

// conver map to json
func (p ProductsQuery) mapToJSON(m interface{}) (string, bool) {
	// var jsonData []byte
	mapData, ok := m.(map[string]interface{})
	jsonData_ := ""
	is_ok := false

	if ok {
		jsonData, err := json.Marshal(mapData)
		if err == nil {
			is_ok = true
			jsonData_ = string(jsonData)
		}
	}

	return jsonData_, is_ok
}