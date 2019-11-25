package goproxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ClientID int64  `json:"clientId"`
}

type loginRes struct {
	Valid bool `json:"valid"`
	Code  string  `json:"code"`
}

func TestGoProxy_Do(t *testing.T) {
	var gp GoProxy
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

		fmt.Println("uRes: ", uRes)
		//fmt.Println("ur: ", ur)
		if !suc || resp.StatusCode != 200 || uRes.Valid == false || uRes.Code != "1" {
			t.Fail()
		}
	}
}