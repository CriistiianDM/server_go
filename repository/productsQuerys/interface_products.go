/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
*/
package productsQuerys

// Librerary import

/*  Interface declaration */
type InterfaceProductsQuery interface {
	InserProductCompanys(params map[string]interface{}) (bool, error)
}

// Return InserProductCompanys
func InserProductCompanys(p InterfaceProductsQuery, params map[string]interface{}) (bool, error) {
	return p.InserProductCompanys(params)
}