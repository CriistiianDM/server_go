/**
   * 
   * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
   * @copyrigth: 2023
   * @license: GPL-3.0
*/
package httpControl

import (
	"fmt"
  "github.com/gin-gonic/gin"
  "container/list"
)

/**
  *
  * General Control for HTTP Request
*/
type HttpRequestControl struct {
  ParamsRequest *list.List
  BodyRequest *gin.Context
}

func (p *HttpRequestControl) GetParamsRequest() *list.List {
  return p.ParamsRequest
}

func (p HttpRequestControl) GetBodyRequest() []byte {
    // Initialize body request
    _body := []byte{}
    // Validate if body request is empty
    if p.BodyRequest != nil {

        // Get body request
        c := p.BodyRequest

        // Get body request
        body, err := c.GetRawData()

        // Validate if body request is not empty
        if err == nil {
          _body = body 
        } else { fmt.Println("No can get body request", err) }

    }
    return _body
}