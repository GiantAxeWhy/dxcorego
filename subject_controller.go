package main

import (
	"fmt"
	"github.com/GiantAxeWhy/dxcorego/framework/gin"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")

}

func SubjectListController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")

}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")

}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")

}

func SubjectGetController(c *gin.Context) {
	subjectId, _ := c.DefaultParamInt("id", 0)
	c.ISetOkStatus().IJson("ok, SubjectGetController:" + fmt.Sprint(subjectId))

}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")

}
