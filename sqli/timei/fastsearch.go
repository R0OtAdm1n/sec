package timei

import (
	"fmt"
	"sqli/timei/fun"
	"sqli/utils"
	"sqli/utils/http"
	"sync"
)

func Fastsearch()  {
	var flag [32]string
	wg := sync.WaitGroup{}
	httpClient := http.HttpClient()
	for i:=0;i<32;i++{
		wg.Add(1)
		go func(i int) {
			str := fun.Binarysearch(i+1,httpClient,utils.Args)
			flag[i]=str
			fmt.Printf("%v ===>>> %v\n",i+1,str)
			wg.Done()
		}(i)
	}
	wg.Wait()
	res:=""
	for _,v:=range flag {
		res+=v
	}
	fmt.Printf("Get Flang ===>>> %v\n",res)
}