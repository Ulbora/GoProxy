//Package goproxy ...
package goproxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ClientID int64  `json:"clientId"`
}

type loginRes struct {
	Valid bool   `json:"valid"`
	Code  string `json:"code"`
}

func TestGoProxy_Do(t *testing.T) {
	var gp GoProxy
	p := gp.GetNewProxy()
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
		var uRes loginRes
		suc, stat := p.Do(req, &uRes)

		fmt.Println("suc: ", suc)

		fmt.Println("uRes: ", uRes)

		if !suc || stat != 200 || uRes.Valid == false || uRes.Code != "1" {
			t.Fail()
		}
	}
}

func TestGoProxy_DoFail(t *testing.T) {
	var gp GoProxy
	p := gp.GetNewProxy()
	var ulogin = new(login)
	ulogin.Username = "ken"
	ulogin.Password = "ken"
	ulogin.ClientID = 1
	var sURL = "http://localhost:3002/rs/user/login"
	aJSON, err := json.Marshal(ulogin)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", sURL, bytes.NewBuffer(aJSON))
		req.Header.Set("Content-Type", "application/json")
		if rErr != nil {
			fmt.Println("rErr: ", rErr)
		}

		var uRes loginRes

		suc, stat := p.Do(req, &uRes)

		fmt.Println("suc: ", suc)

		fmt.Println("uRes: ", uRes)
		fmt.Println("stat: ", stat)
		if suc {
			t.Fail()
		}
	}
}

func TestGoProxy_DoFail2(t *testing.T) {
	var gp GoProxy
	p := gp.GetNewProxy()
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
		var uRes loginRes
		suc, stat := p.Do(req, nil)

		fmt.Println("suc: ", suc)

		fmt.Println("uRes: ", uRes)
		fmt.Println("stat: ", stat)

		if suc || stat != 200 || uRes.Valid != false {
			t.Fail()
		}
	}
}
