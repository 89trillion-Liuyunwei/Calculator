package api

import (
	"github.com/gin-gonic/gin"
	"max/model"
)

func Calculation(c *gin.Context) {
	data := c.PostForm("data")        //接受计算的表达式data参数
	code := 0                         //初始化响应编号
	message := ""                     //初始化响应信息
	re := 0                           //初始化计算结果
	if data == "" && len(data) == 0 { //数据为空返回提示信息
		code = 1
		message = "请不要提交空数据"
	} else {
		num, errors := model.Calculation(data) //进行计算
		if errors == nil {                     //计算成功返回计算结果
			code = 0
			re = num
			message = "success"
		} else { //计算错误返回错误结果
			code = 1
			message = errors.Error()
		}
	}
	c.JSON(200, gin.H{
		"code":    code,
		"data":    re,
		"message": message,
	})
}
