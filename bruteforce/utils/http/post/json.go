package post

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"sqli/utils/http"
)

/**
发生json请求
 */
func Json(args map[string]interface{}) *fasthttp.Request {
	req := http.Request()
	marshal,_:=json.Marshal(args)
	req.SetBody(marshal)
	return req
}