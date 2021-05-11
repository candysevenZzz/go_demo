package main

import (
	"github.com/gin-gonic/gin"
	"pratice/src/common"
	"pratice/src/router"
)

func main() {
	//指定服务模式
	gin.SetMode(gin.DebugMode)

	// 创建一个默认的路由引擎
	engine := gin.Default()

	//加载路由
	router.InitRouter(engine)

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	_ = engine.Run(common.Addr + ":" + common.Port)
}
