package goproxy

import "net/http"

//Proxy Proxy
type Proxy interface {	
	Do(req *http.Request) (bool, *http.Response)
}
