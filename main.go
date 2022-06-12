package main

import (
	"goweb/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	ser := http.Server{
		Handler: core,
		Addr:    ":8080",
	}
	ser.ListenAndServe()
}
