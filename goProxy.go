package goproxy

import (
	"fmt"
	"net/http"
)

//GoProxy GoProxy
type GoProxy struct {
}


//Do Do
func (p *GoProxy) Do(req *http.Request) (bool, *http.Response)  {
	var suc bool
	// var status int
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("Do err: ")
		fmt.Println(err)
	}else{
		suc = true
	}	
	return suc, resp
}

//go mod init github.com/Ulbora/GoProxy
