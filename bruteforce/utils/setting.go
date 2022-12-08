package utils

import (
	"github.com/valyala/fasthttp"
	"log"
	"github.com/spf13/viper"
)

var (
	Userpath string
	Passpath string
	Goroutine int
	//Timeout int
	Url string
	Method string
	Header map[string]string
	// Body map[string]string
	Args fasthttp.Args
)

func Init()  {
	config := viper.New()
	config.AddConfigPath("D:\\MyProgram\\sec\\sqli\\config")
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("找不到配置文件..")
		} else {
			log.Println("配置文件出错..")
		}
	}
	Load(config)
}

func Load(config *viper.Viper)  {
	Userpath = config.GetString("userPath")
	Passpath = config.GetString("passPath")
	Goroutine = config.GetInt("goroutine")
	//Timeout = config.GetInt("timeout")
	Url = config.GetString("url")
	Method = config.GetString("method")
	Header = config.GetStringMapString("header")
	Body := config.GetStringMapString("body")

	Args := fasthttp.AcquireArgs()

	for k,v := range Body {
		Args.Set(k,v)
	}
}