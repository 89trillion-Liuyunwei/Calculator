package model

import (
	"errors"
	"strconv"
)

type maths interface {
	add(x int, y int) int                      //加法
	sub(x int, y int) int                      //减法
	mul(x int, y int) int                     //乘法
	division(x int, y int) int                 //除法
	modulo(x int, y int)                       //取余
	pow(x int, y int)                          //幂函数
	Operator(x int, y int, sign string) string //计算
	JudgmentType(s string, i int) bool         //判断一个字节是不是0-9的数字
	IsNum(s string) bool                       //判断是不是数字
	isLower(top string, newTop string) bool    //比较符号优先级

}
type Math struct {
}

func (math Math) add(x int, y int) int {
	return x + y
}
func (math Math) sub(x int, y int) int {
	return x - y
}
func (math Math) mul(x int, y int) int {
	return x * y
}
func (math Math) division(x int, y int) int {
	return x / y
}
func (math Math) modulo(x int, y int) int {
	return x % y
}
func (math Math) pow(x int, y int) int {
	pow := 1
	for j := 1; j <= y; j++ {
		pow = pow * x
	}
	return pow
}

/*
  JudgmentType  判断字节是否为数字函数
 *判断字符串中的一个字节是否是数字
 *请求参数:s(string)预判断的字符串  i(int)预判断字节在字符串中的位置
 *正确范围:数字0-9 && 预判断i未越界
 *返回结果：true(合格) false(错误)
*/
func (math Math) JudgmentType(s string, i int) bool {
	if i > len(s)-1 {
		return false
	}
	if s[i]-'0' >= 0 && s[i]-'0' <= 9 { //判断是否为0-9内数字
		return true
	}
	return false
}

//判断是否是一个数字
func (math Math) IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// 比较运算符栈栈顶 top 和新运算符 newTop 的优先级高低 newTop优先级高则返回true
func (math Math) IsLower(top string, newTop string) bool {
	switch top {
	case "+", "-": //+ -为一级优先级
		if newTop == "*" || newTop == "/" || newTop == "%" || newTop == "^" {
			return true
		}
	case "*", "/", "%": //* / 为二级级优先级
		if newTop == "^" {
			return true // ^是三级优先级
		}
	case "(": // （  最高优先级
		return true
	}
	return false
}

/*
  Operator  基本计算
 *输入计算符号和数字进行计算
 *请求参数:x(string)计算符号左边数  y(int)计算符号右边数   nexChar(string) 计算符号
 *正确范围:如果是除法，除数不能为0
 *返回结果：计算结果，错误提示
*/
func (math Math) Operator(x int, y int, nexChar string) (string, error) {
	num := 0
	var _ Math
	switch nexChar {
	case "+":
		num = math.add(x, y)
	case "-":
		num = math.sub(x, y)
	case "*":
		num = math.mul(x, y)
	case "/":
		if y == 0 {
			return "", errors.New("除数不能为0")
		}
		num = math.division(x, y)
	case "%":
		num = math.modulo(x, y)
	case "^":
		num = math.pow(x, y)
	}
	return strconv.Itoa(num), nil

}
