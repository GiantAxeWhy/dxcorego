package main

import (
	"dxcorego/framework"
	"dxcorego/framework/middleware"
)

func registerRouter(core *framework.Core) {
	// core.Get("foo", framework.TimeoutHandler(FooControllerHandler, time.Second*1))
	//静态路由+HTTP方法匹配
	core.Get("/user/login", middleware.Test3(), UserLoginController)
	//批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test3())
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}