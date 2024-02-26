/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package defaultDomain

/** Interface */
type InterfaceDefaultDomain interface {
	DomGetData(p *map[string]interface{})
}

/**
  * Return all allowed ips
*/
func DomGetData(p InterfaceDefaultDomain , args *map[string]interface{}) {
	p.DomGetData(args)
}