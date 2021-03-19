package model

import (
	"errors"
	"strconv"
)

/*

从左到右逐个字符遍历中缀表达式，输出的字符序列即是后缀表达式：
循环：
	判断字节是否符合规则，不符合抛出错误
	遇到数字直接输出到postfix切片里面
	遇到运算符则判断：
		如果是空格跳过
		栈顶运算符优先级更低则入栈，更高或相等则直接输出到postfix切片里面
		是 ( 直接入栈
		运算符是 ) 则将栈顶运算符全部弹出，直到遇见 (
	中缀表达式遍历完毕，运算符栈不为空则全部弹出，依次追加到输出
例子：
	中缀表达式:5-10*(2/1)^2
	后缀表达式:5 10 2 1 /  2 ^ * -

*/
func Calculation(s string) (int, error) {
	var myStack *Stack = new(Stack)
	var postfix = make([]string, 100)
	var j int = 0
	var math Math
	l := len(s)
	num := 0
	for i := 0; i < l; i++ {
		if !Stringently(s, i) { // 判断字符串中的每个字节是否合法，不合法抛出错误 正确范围:数字0-9 和空格
			return 0, errors.New("预计算的字符串格式错误")
		}
		char := string(s[i])
		if math.JudgmentType(s, i) {
			its := int(s[i] - '0') //uint8转换为int类型
			num = num*10 + its     //取字符串中的一个数字  例如：53可通过2次循环num*10 + ints获取到
			if !math.JudgmentType(s, i+1) {
				postfix[j] = strconv.Itoa(num)
				num = 0
				j++
			}
		} else {
			switch char {
			case " ":
				continue
			case "(": // 左括号直接入栈
				myStack.Push("(")
			case ")": // 右括号则弹出元素直到遇到左括号
				{
					for ; !myStack.IsEmpty(); j++ {
						preChar, err := myStack.Pop()
						if err != nil {
							return 0, err
						}
						if preChar == "(" {
							break
						}
						postfix[j] = preChar
					}

				}
			default: // 遇到高优先级的运算符，不断弹出，直到遇见更低优先级运算符或者"("
				for ; !myStack.IsEmpty(); j++ {
					top, err := myStack.Top()
					if err != nil {
						return 0, err
					}
					if top == "(" || math.IsLower(top, char) {
						break
					}
					postfix[j] = top
					_, err = myStack.Pop()
					if err != nil {
						return 0, err
					}
				}
				myStack.Push(char) // 低优先级的运算符入栈
			}
		}
	}
	for ; !myStack.IsEmpty(); j++ { //运算符栈不为空则全部弹出，依次追加到输出
		temp, err := myStack.Pop()
		if err != nil {
			return 0, err
		}
		postfix[j] = temp
	}
	re, err := postfixCuration(postfix) //后缀表达式计算
	if err != nil {
		return 0, err
	}
	return re, nil
}

/*

从左到右逐个字符遍历后缀表达式，输出的是计算结果：
循环：
	如果是数字压入栈
	如果是运算符，输出栈内两个元素进行计算，将结果压入栈
	遍历完返回计算结果

*/

func postfixCuration(postfix []string) (int, error) {
	var stack *Stack = new(Stack)
	var math Math
	fixLen := len(postfix)
	for i := 0; i < fixLen; i++ {
		nextChar := string(postfix[i])
		if math.IsNum(postfix[i]) {
			stack.Push(nextChar) // 数字：直接压栈
		} else if nextChar != "" { // 操作符：取出两个数字计算值，再将结果压栈
			num1, err := stack.Pop()
			if err != nil {
				return 0, errors.New("预计算的字符串格式错误")
			}
			num2, err := stack.Pop()
			if err != nil {
				return 0, errors.New("预计算的字符串格式错误")
			}
			num3, err := strconv.Atoi(num1)
			num4, err := strconv.Atoi(num2)
			temp, err := math.Operator(num4, num3, nextChar)
			if err != nil {
				return 0, err
			}
			stack.Push(temp) //进行符号计算
		} else {
			break
		}
	}
	re, err := stack.Top()
	if err != nil {
		return 0, err
	}
	result, _ := strconv.Atoi(re)
	return result, nil
}
