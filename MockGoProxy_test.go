package goproxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestMockGoProxy_Do(t *testing.T) {
	var bdy loginRes
	bdy.Valid = true
	bdy.Code = "1"
	var res http.Response
	res.StatusCode = 200
	res.Body = ioutil.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))
	var gp MockGoProxy
	gp.MockDoSuccess1 = true
	gp.MockResp = &res
	//gp.Do()
	var p Proxy
	p = &gp
	var ulogin = new(login)
	ulogin.Username = "ken"
	ulogin.Password = "ken"
	ulogin.ClientID = 1
	var sURL = "http://localhost:3001/rs/user/login"
	aJSON, err := json.Marshal(ulogin)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", sURL, bytes.NewBuffer(aJSON))
		req.Header.Set("Content-Type", "application/json")
		if rErr != nil {
			fmt.Println("rErr: ", rErr)
		}
		//var ur interface{}
		var uRes loginRes
		//ur = uRes
		suc, resp := p.Do(req)
		if suc {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&uRes)
			if error != nil {
				log.Println(error.Error())
			}
		}
		fmt.Println("suc: ", suc)
		fmt.Println("res: ", resp)

		fmt.Println("uRes mock: ", uRes)
		fmt.Println("resp mock: ", resp)
		if !suc || resp.StatusCode != 200 || uRes.Valid == false || uRes.Code != "1" {
			t.Fail()
		}
	}
}
