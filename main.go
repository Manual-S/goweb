package main

import (
	"goweb/framework/gin"
	"goweb/framework/provider/app"
	"goweb/provider/demo"
	"net/http"
)

func main() {
	core := gin.New()
	core.Bind(&demo.DemoServiceProvider{})
	core.Bind(&app.DirectoryProvider{})
	// 注册路由
	registerRouter(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8080", // 监听本机的8080端口
	}

	server.ListenAndServe()
}
