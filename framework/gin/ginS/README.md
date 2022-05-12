# Gin Default Server

This is API experiment for Gin.

```go
package main

import (
	"github.com/GiantAxeWhy/dxcorego/framework/gin"
	"github.com/GiantAxeWhy/dxcorego/framework/gin/ginS"
)

func main() {
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
```
