package model

const filter string = "+-*/^%1234567890 ()" //符号白名单
/*
  Stringently 判断字节是否合格函数
 *判断字符串中的一个字节是否符合计算基本要求
 *请求参数:s(string)预判断的字符串  i(int)预判断字节在字符串中的位置
 *正确范围:数字0-9 和空格
 *返回结果：true(合格) false(错误)
*/
func Stringently(s string, i int) bool {
	if i > len(s)-1 {
		return true
	}
	for j := 0; j < len(filter); j++ {
		if s[i] == filter[j] {
			return true
		}
	}
	return false
}
