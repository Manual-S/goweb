package main

import (
	"goweb/app/console"
	"goweb/app/http"
	"goweb/framework"
	"goweb/framework/provider/app"
	"goweb/framework/provider/distributed"
	"goweb/framework/provider/env"
	"goweb/framework/provider/kernel"
	"log"
)

func main() {

	// 设置日志输出选项
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)

	container := framework.NewContainer()
	container.Bind(&app.DirectoryProvider{})               // 绑定目录结构服务
	container.Bind(&distributed.LocalDistributeProvider{}) // 绑定分布式定时器
	container.Bind(&env.EnvProvider{})

	// 绑定一个路由服务
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.KernelProvider{HttpEngine: engine})
	}

	console.RunCommand(container)
}
