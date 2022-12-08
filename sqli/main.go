package main

import (
	"fmt"
	"sqli/timei"
	"sqli/utils"
	"time"
)

func main()  {
	start:=time.Now()
	utils.Init()
	if utils.Goroutine {
		timei.Fastsearch()
	} else {
		timei.Search()
	}
	end:=time.Since(start)
	fmt.Printf("Cost %v",end)
}