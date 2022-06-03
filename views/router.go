package views

import (
	"github.com/gin-gonic/gin"
	"mydns/helper"
	"mydns/logs"
	"mydns/views/hello"
)

func routes(r *gin.Engine) {
	dparam := helper.DefaultOptionalParam()
	dparam.CheckSign = true

	r.GET("/hello/", helper.Api(hello.Hello, dparam))
	r.GET("/hello/error/", helper.Api(hello.TestError, dparam))
	r.GET("/hello/panic/", helper.Api(hello.TestPanic, dparam))
}

func InitRoutes(r *gin.Engine) {
	routes(r)

	r.NoRoute(func(c *gin.Context) {
		logs.Warn("url doesn't match, uri=%v", c.Request.RequestURI)
		c.JSON(200, gin.H{
			"status_code": 0,
			"status_msg":  "url doesn't match",
		})
	})

}
