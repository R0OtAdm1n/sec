package http

import (
	"github.com/valyala/fasthttp"
	"sqli/utils"
)

func Request() *fasthttp.Request {
	req:=fasthttp.AcquireRequest()
	req.Header.SetMethod(utils.Method)
	req.SetRequestURI(utils.Url)
	for k,v := range utils.Header {
		req.Header.Set(k,v)
	}
	return req
}