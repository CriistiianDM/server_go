/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/
package utils

/**
  * Validate if propierties of file .env have info
*/
func isNotEmptyEnv(args ...string) bool {
	isEmpty := false

	if len(args) > 0 {
		for _, env := range args {
			if env == "" {
				isEmpty = false
				break
			} else {isEmpty = true}
		}
	}
	
	return isEmpty
}

// Check keys of map
func CheckExpectedKeys(m map[string]interface{}, expectedKeys []string) bool {
	haveAllKeys := true;
    for _, key := range expectedKeys {
        if _, ok := m[key]; !ok {
            haveAllKeys =  false
        }
    }

    return haveAllKeys;
}