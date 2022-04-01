package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"apiserver/config"
	"apiserver/model"
	"apiserver/pkg/logger"
	"apiserver/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file")
)

func main() {

	pflag.Parse()

	//init config
	//在 main 函数中增加了 config.Init(*cfg) 调用，用来初始化配置，cfg 变量值从命令行 flag 传入，可以传值，比如 ./apiserver -c config.yaml，也可以为空，如果为空会默认读取 conf/config.yaml
	if err := config.Init(*cfg); err != nil {
		errors.New(err.Error())
	}

	//init logger config
	//logger.Init()

	//init DB
	model.DB.Init()

	// Set gin mode.//gin 有 3 种运行模式：debug、release 和 test，其中 debug 模式会打印很多 debug 信息。
	gin.SetMode(viper.GetString("runmode"))

	//Create the Gin engine
	g := gin.New()

	//gin middlewares
	var middlewares []gin.HandlerFunc

	//Routes
	router.Land(
		// Cores
		g,

		// Middlewares
		middlewares...,
	)

	logger.Info("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	logger.Error(http.ListenAndServe(viper.GetString("addr"), g).Error())

	go func() {
		if err := pingServer(); err != nil {
			logger.Exception("The router has no response, or it might took too long to start up.", err)

		}
		logger.Info("The router has been deployed successfully.")
	}()
}

func pingServer() error {

	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
