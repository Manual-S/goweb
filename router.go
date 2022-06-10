package main

import "goweb/framework"

func registerRouter(core *framework.Core) {
	core.Get("/user/login", FooControllerHandler)

	subjectApi := core.Group("/subject")
	{
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}
