package main

import (
	"goweb/framework/gin"
)

func registerRouter(core *gin.Engine) {
	core.GET("subject/list/all", SubjectListController)
}
