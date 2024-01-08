/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
*/

package query

import (
	"server_go/db"
	"reflect"
	"database/sql"
	"fmt"
)
var dbInstance = db.GetConnection()

/**
  * Define querys
*/
type GeneralQuery struct {}

/**
  * Return all params in request
*/
func (p *GeneralQuery) Query(sqlQuery string, args ...interface{}) (*sql.Rows, error) {
	//return p.queryGeneral("Query", sqlQuery ,args...)
	result, err := p.queryGeneral("Query", sqlQuery, args...)
    if err != nil {
        return nil, err
    }

    // Asegúrate de realizar una aserción de tipo adecuada
    rows, ok := result.(*sql.Rows)
    if !ok {
        return nil, fmt.Errorf("El resultado no es un *sql.Rows")
    }

    return rows, nil
}

/**
  * Execute query
*/
func (p *GeneralQuery) Exec(sqlQuery string, args ...interface{}) (sql.Result, error) {
	//return p.queryGeneral("Exec", sqlQuery ,args...)
	result, err := p.queryGeneral("Exec", sqlQuery, args...)
    if err != nil {
        return nil, err
    }

    // Asegúrate de realizar una aserción de tipo adecuada
    execResult, ok := result.(sql.Result)
    if !ok {
        return nil, fmt.Errorf("El resultado no es un sql.Result")
    }

    return execResult, nil
}

/**
  * Estrucuture query
*/
func (p *GeneralQuery) queryGeneral(methodName string, sqlQuery string ,args ...interface{}) (interface{}, error) {
	var ( 
		result interface{}
		Error error
	)

	if sqlQuery != "" {
		dbValue := reflect.ValueOf(dbInstance)

		method := dbValue.MethodByName(methodName)
		if method.IsValid() {

			argValues := make([]reflect.Value, len(args)+1)
			argValues[0] = reflect.ValueOf(sqlQuery)
			for i, arg := range args {
				argValues[i+1] = reflect.ValueOf(arg)
			}

			resultValues := method.Call(argValues)

			result = resultValues[0].Interface().(*sql.Row)

		}
	}
	return result, Error
}