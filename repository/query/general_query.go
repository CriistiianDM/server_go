/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
*/

package query

import (
	"server_go/db"
	_ "reflect"
	"database/sql"
	_ "fmt"
)
var dbInstance = db.GetConnection()

/**
  * Define querys
*/
type GeneralQuery struct {}

/**
  * Return all params in request
*/
func (p GeneralQuery) Query(sqlQuery string, args ...interface{}) (*sql.Rows, error) {
	dbInstance = db.GetConnection()
	return dbInstance.Query(sqlQuery, args...)
}

/**
  * Execute query
*/
func (p GeneralQuery) Exec(sqlQuery string, args ...interface{}) (sql.Result, error) {
	dbInstance = db.GetConnection()
	return dbInstance.Exec(sqlQuery, args...)
}

/**
  * Estrucuture query
*/
// func (p *GeneralQuery) queryGeneral(methodName string, sqlQuery string ,args ...interface{}) (interface{}, error) {
// 	var ( 
// 		result = new(sql.Row)
// 		Error error
// 	)
// 	dbInstance = db.GetConnection()
// 	if sqlQuery != "" {
// 		dbValue := reflect.ValueOf(dbInstance)

// 		method := dbValue.MethodByName(methodName)
// 		if method.IsValid() {

// 			argValues := make([]reflect.Value, len(args)+1)
// 			argValues[0] = reflect.ValueOf(sqlQuery)
// 			for i, arg := range args {
// 				argValues[i+1] = reflect.ValueOf(arg)
// 			}

// 			resultValues := method.Call(argValues)

// 			result_ := resultValues[0].Interface().(*sql.Row)
// 			fmt.Println("Resulttt", result_)
// 		}
// 	}
// 	return result, Error
// }