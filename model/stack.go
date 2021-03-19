package model

import "errors"

type Stack []string

/*
栈数据结构定义
*/
//获得栈长度
func (stack Stack) Len() int {
	return len(stack)
}

//判断栈是不是为空
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

//获得栈容量
func (stack Stack) Cap() int {
	return cap(stack)
}

//入栈
func (stack *Stack) Push(value string) {
	*stack = append(*stack, value)
}

//获得栈顶元素
func (stack Stack) Top() (string, error) {
	if len(stack) == 0 {
		return "", errors.New("获得失败，栈为空")
	}
	return stack[len(stack)-1], nil
}

//出栈
func (stack *Stack) Pop() (string, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return "", errors.New("出栈失败，栈为空")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}
