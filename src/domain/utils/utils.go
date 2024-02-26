/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/

package utils;

import (
	"github.com/gin-gonic/gin"
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
	"encoding/json"
    "io"
)

var (
	keyToken = "12378chhcbcg673gfsvgh654"
)

/**
  * Conver data of request to []*string
  * Todo: Add Validation later
*/
func ConvertDataToArrayString(p map[string]interface{}) []*string {
	var response []*string
	expectedKeys := []string{"email", "password", "nickname", "firstname", "lastname", "img"}
	if len(p) == 6 && CheckExpectedKeys(p,expectedKeys) {
	   for _, key := range expectedKeys {
		 value := p[key].(string)
		 response = append(response, &value)
	   }
	}

	return response;
}

/**
  * Conver data of request to []*int
  * Todo: Add Validation later
*/
func ConvertDataToArrayInt(p map[string]interface{}) []*int {
	var response []*int
	expectedKeys := []string{"user_account", "follower_id"}

	if CheckExpectedKeys(p,expectedKeys) {
	   for _, key := range expectedKeys {
		 if value, ok := p[key].(float64); ok {
			valueInt := int(value)
		 	response = append(response, &valueInt)
		 } 
	   }
	}
	
	return response;
}

/**
  * Conver data of request to []*interface{} 
  * Todo: Add Validation later
*/
func ConvertDataToArrayInterface(p map[string]interface{}) []*interface{} {
	var response []*interface{}
	expectedKeys, ok := p["expectedKeys"].([]string)

	if ok && CheckExpectedKeys(p,expectedKeys) {
	   for _, key := range expectedKeys {
		 if value, ok := p[key]; ok {
		 	response = append(response, &value)
		 } 
	   }
	}
	
	return response;
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

/**
  * Return res to client
*/
func ResponseControlGeneral(c *gin.Context , args ...map[string]interface{}) {
	// Instance data of the request
	response := args[0]
	_status := response["status"].(int)
	_key := []byte(keyToken)
	_res := gin.H{
		"data": response["data"],
		"status": response["statusReq"], 
		"message": response["message_default"],
	}
	jsonRes, err := json.Marshal(_res)
	
	if err ==  nil {
		encryptedRes, err := encrypt(jsonRes, _key)
		if err == nil {
			encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedRes)
			_res = gin.H{"res": encryptedBase64}
		}
	}

	c.JSON(_status, _res);
}

/**
  * Data default to res 
*/
func StateDefaultReq() map[string]interface{} {
	return map[string]interface{}{
		"data": nil,
		"statusReq": false, 
		"message_default": "Error in request",	
		"status": 500,
	}
}

/* encrypt AES */
func encrypt(data []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
	var ciphertext []byte

    if err == nil {

		ciphertext = make([]byte, aes.BlockSize + len(data))
		iv := ciphertext[:aes.BlockSize]
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return nil, err
		}
	
		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
    }

    return ciphertext, nil
}