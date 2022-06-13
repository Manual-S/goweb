package main

import (
	"goweb/framework/gin"
	"net/http"
)

func main() {
	eng := gin.New()
	Router(eng)
	server := http.Server{
		Handler: eng,
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
