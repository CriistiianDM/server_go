/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package allowedIPDomain

/** Interface */
type InterfaceAlloIP interface {
	DomGetData(p *map[string]interface{})
}

/**
  * Return all allowed ips
*/
func DomGetData(p InterfaceAlloIP , args *map[string]interface{}) {
	p.DomGetData(args)
}