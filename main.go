package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"max/route"
	"os"
)


func main() {
	//gin.DisableConsoleColor()
	//创建日志
	f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	//初始化路由
	r := route.InitRoute()
	r.Run(":8080")
}

//1+2*3+5^2%3
