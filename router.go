package main

import (
	"goweb/framework"
)

func registerRouter(core *framework.Core) {
	//core.Use(middleware.Test1())
	//core.Use(middleware.Test2())
	//core.Use(middleware.Test1())
	core.Get("/user/login", FooControllerHandler)
	core.Get("/subject/:id", SubjectGetController)
}
