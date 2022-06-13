package main

import (
	"goweb/framework/gin"
)

func HelloFunc(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("ok")
}
