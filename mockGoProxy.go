package goproxy

import (
	"encoding/json"
	"log"
	"net/http"
)

//MockGoProxy MockGoProxy
type MockGoProxy struct {
	MockDoSuccess1 bool
	MockRespCode   int
	MockResp       *http.Response
}

//Do Do
func (p *MockGoProxy) Do(req *http.Request, obj interface{}) (bool, int) {
	defer p.MockResp.Body.Close()
	decoder := json.NewDecoder(p.MockResp.Body)
	error := decoder.Decode(obj)
	if error != nil {
		log.Println("Decode Error in Mock: ", error.Error())
	}
	return p.MockDoSuccess1, p.MockRespCode
}

//GetNewProxy GetNewProxy
func (p *MockGoProxy) GetNewProxy() Proxy {
	var px Proxy
	px = p
	return px
}
