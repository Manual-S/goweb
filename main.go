package main

import (
	"context"
	"goweb/framework"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	ser := http.Server{
		Handler: core,
		Addr:    ":8080",
	}
	go func() {
		ser.ListenAndServe()
	}()

	// 监听进程终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 没有信号到来 主goroutine阻塞
	<-quit

	if err := ser.Shutdown(context.Background()); err != nil {
		log.Fatalf("server Shutdown error %v", err)
	}
}
