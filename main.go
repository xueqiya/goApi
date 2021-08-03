package main

import (
	"github.com/gin-gonic/gin"
	"xueqiya/config"
	"xueqiya/config/route"
)

func main() {
	new(config.Database).GetConnect()
	defer config.DB.Close()
	engine := gin.Default()
	(&route.Route{Engine: engine}).Run()
	engine.Run(":8080")
}
