package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pratice/src/common"
	v1 "pratice/src/controller/v1"
	v2 "pratice/src/controller/v2"
	"strconv"
)

func InitRouter(engine *gin.Engine) {
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	engine.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(common.HttpCode200, gin.H{
			"message": "Hello world!",
		})
	})

	engine.GET("/sn", SignDemo)

	// v1 版本
	GroupV1 := engine.Group("/v1")
	{
		GroupV1.Any("/product/add", v1.AddProduct)
		GroupV1.Any("/member/add", v1.AddMember)
	}

	// v2 版本
	GroupV2 := engine.Group("/v2")
	{
		GroupV2.Any("/product/add", v2.AddProduct)
		GroupV2.Any("/member/add", v2.AddMember)
	}
}

func SignDemo(c *gin.Context) {
	ts := strconv.FormatInt(common.GetTimeUnix(), 10)
	fmt.Println(ts)
}
