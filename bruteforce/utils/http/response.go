package http

import (
	"github.com/valyala/fasthttp"
)

//var httpClient = httpclient.HttpClient()

// 获取http响应报文
func Response(httpClient *fasthttp.Client,req *fasthttp.Request) (*fasthttp.Response,error) {
	//log.Println(req)
	res:=fasthttp.AcquireResponse()
	err := httpClient.Do(req,res)
	//log.Println(res)
	return res,err
}

/**
func Response(req *fasthttp.Request) (*fasthttp.Response,error) {
	res:=fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(res)
		fasthttp.ReleaseRequest(req)
	}()
	//httpClient := httpclient.HttpClient()
	err := httpClient.Do(req,res)
	if err!=nil {
		log.Println("Response Error!!!")
		return nil,err
	}
	return res,nil
}*/