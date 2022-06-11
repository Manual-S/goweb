package main

import (
	"goweb/framework"
	"goweb/framework/middleware"
)

func registerRouter(core *framework.Core) {
	core.Use(middleware.Test1())
	core.Use(middleware.Test2())
	core.Use(middleware.Test1())
	core.Get("/user/login", FooControllerHandler)
	core.Get("/subject/:id", SubjectListController)
}
