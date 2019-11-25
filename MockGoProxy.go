package goproxy

import "net/http"

//MockGoProxy MockGoProxy
type MockGoProxy struct {
	MockDoSuccess1 bool
	MockResp       *http.Response
}

//Do Do
func (p *MockGoProxy) Do(req *http.Request) (bool, *http.Response) {
	return p.MockDoSuccess1, p.MockResp
}
