package main

import (
	"flag"
	"time"
	"audit-rule/config"
	"github.com/kataras/iris"
	"audit-rule/service"
	"github.com/iris-contrib/graceful"
)

func main() {
	conf := flag.String("conf", "../config/app.conf", "config file path")
	listen := flag.String("listen", ":8088", "listen address and port")
	flag.Parse()
	config.Init(*conf)
	api := iris.New()
	api.Use()
	srv := service.NewService()
	srv.RegisterAll(api)
	graceful.Run(*listen, time.Duration(10)*time.Second, api)
}
