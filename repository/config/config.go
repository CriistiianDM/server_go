/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package config;

/**
  * Config Map Sql
*/
func ConfigMapSql() map[string]string {
	return map[string]string{
		"default_routes": "all_routes",
		"get_allowed_ips": "allowed_ips",
	}
}