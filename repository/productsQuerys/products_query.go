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
	_insert_query = "CALL insert_new_product_per_company($1,$2)"
	_get_query = "SELECT * FROM products_per_company($1)"
)

type ProductsQuery struct {}


/** Aux Class */
type ParamsCompanyProducts struct {
	id int
    route string
    json_data string
    created_at string
    updated_at string
}

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
		result, err := query.Exec(query.GeneralQuery{},_insert_query, company, json_)
		is_err = err
		if err == nil {
			fmt.Println("Result", result)
			response = !response
		}
	}

	return response, is_err
}

func (p ProductsQuery) GetProductsPerCompany(params map[string]interface{}) (map[string]interface{} , error) {
	company, ok_company :=  params["company"];
	response := make(map[string]interface{})
	var is_err error
	
	if ok_company {
		result, err := query.Query(query.GeneralQuery{},_get_query, company);

		if err == nil {
			defer result.Close()
			for result.Next() {
				instance := ParamsCompanyProducts{}
				err := result.Scan(&instance.id, &instance.route ,&instance.json_data ,&instance.created_at, &instance.updated_at)
				if err != nil {
					fmt.Println("Error", err)
				}
				response[instance.route] = instance
			}
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