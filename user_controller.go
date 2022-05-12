package main

import (
	"github.com/GiantAxeWhy/dxcorego/framework/gin"
	"time"
)

func UserLoginController(c *gin.Context) {
	foo, _ := c.DefaultQueryString("foo", "def")
	// 等待10s才结束执行
	time.Sleep(10 * time.Second)
	c.ISetOkStatus().IJson("ok, UserLoginController: " + foo)

}
