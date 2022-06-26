package main

import (
	"fmt"
	"goweb/app/console"
	"goweb/app/http"
	"goweb/framework"
	"goweb/framework/provider/app"
	"goweb/framework/provider/kernel"
)

func main() {
	container := framework.NewContainer()
	container.Bind(&app.DirectoryProvider{}) // 绑定目录结构服务

	// 绑定一个路由服务
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.KernelProvider{HttpEngine: engine})
		fmt.Printf("bind engine succ\n")
	}

	console.RunCommand(container)
}
