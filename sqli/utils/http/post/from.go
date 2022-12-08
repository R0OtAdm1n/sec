package post

import (
	"github.com/valyala/fasthttp"
	"sqli/utils/http"
)

// 构造POST请求的http报文
func Form(args fasthttp.Args) *fasthttp.Request {
	req := http.Request()
	args.WriteTo(req.BodyWriter())
	//log.Println(req)
	return req
}