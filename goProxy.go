package goproxy

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//GoProxy GoProxy
type GoProxy struct {
}

//Do Do
func (p *GoProxy) Do(req *http.Request, obj interface{}) (bool, int) {
	var suc bool
	var status int
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Do err: ", err)
		fmt.Println("resp in fail: ", resp)
	} else {
		defer resp.Body.Close()
		status = resp.StatusCode
		decoder := json.NewDecoder(resp.Body)
		error := decoder.Decode(obj)
		if error != nil {
			log.Println("Decode Error: ", error.Error())
		} else {
			suc = true
		}
	}
	return suc, status
}

//GetNewProxy GetNewProxy
func (p *GoProxy) GetNewProxy() Proxy {
	var px Proxy
	px = p
	return px
}

//go mod init github.com/Ulbora/GoProxy
