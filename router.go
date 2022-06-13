package main

import "goweb/framework/gin"

func Router(eng *gin.Engine) {
	eng.GET("/hello", HelloFunc)
}
