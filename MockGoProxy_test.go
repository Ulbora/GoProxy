package goproxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMockGoProxy_Do(t *testing.T) {
	var gp MockGoProxy
	var res http.Response
	//res.StatusCode = 200
	res.Body = ioutil.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))
	gp.MockResp = &res
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200

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

		fmt.Println("suc in mock: ", suc)

		fmt.Println("uRes in mock: ", uRes)

		if !suc || stat != 200 || uRes.Valid == false || uRes.Code != "1" {
			t.Fail()
		}
	}
}

func TestMockGoProxy_DoFail(t *testing.T) {
	var gp MockGoProxy
	var res http.Response
	//res.StatusCode = 200
	res.Body = ioutil.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))
	gp.MockResp = &res
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200

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

		fmt.Println("suc in mock fail: ", suc)

		fmt.Println("uRes in mock fail: ", uRes)

		if !suc || stat != 200 || uRes.Valid != false {
			t.Fail()
		}
	}
}
