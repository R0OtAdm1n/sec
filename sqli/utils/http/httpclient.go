package http

import (
	"crypto/tls"
	"github.com/valyala/fasthttp"
	"time"
)

// 创建http连接
func HttpClient() *fasthttp.Client {
	return &fasthttp.Client{
		Name: "HttpClient",
		NoDefaultUserAgentHeader: true,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
		MaxConnsPerHost: 200,
		MaxIdleConnDuration: 5*time.Second,
		MaxConnDuration: 5*time.Second,
		ReadTimeout: 3*time.Second,
		WriteTimeout: 5*time.Second,
		MaxConnWaitTimeout: 5*time.Second,
	}
}
