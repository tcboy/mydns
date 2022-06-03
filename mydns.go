package main

import (
	"github.com/gin-gonic/gin"
	"mydns/logs"
	"mydns/views"
)

func main() {
	r := gin.Default()
	views.InitRoutes(r)

	err := r.Run(":8081")
	if err != nil {
		logs.Error("%s", err.Error())
	}
}
