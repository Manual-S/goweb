// 理解成handler层
package main

import (
	"goweb/framework/gin"
	"goweb/provider/demo"
)

func SubjectListController(c *gin.Context) {
	/*
		这里应该的写法是
		service := NewFooService(c)
		foo := service.GetFoo()
		将foo作为返回值给到前端
	*/
	demoServiceInf := c.MustMake(demo.Key)
	demoService := demoServiceInf.(demo.Service)
	foo := demoService.GetFoo()

	c.ISetOkStatus().IJson(foo)
}
