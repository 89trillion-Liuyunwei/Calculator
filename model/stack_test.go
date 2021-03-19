package model

import (
	"testing"
)

func TestStack_Len(t *testing.T) {
	var stack *Stack = new(Stack)
	*stack = append(*stack, "2")
	*stack = append(*stack, "-")
	*stack = append(*stack, "1")
	var fibTests = []struct {
		expected int // expected result
	}{
		{3},
	}
	for _, tt := range fibTests {
		actual := stack.Len()
		if actual != tt.expected {
			t.Errorf("len() ; actual:%d  ,expected %d", actual, tt.expected)
		}
	}
}
func TestStack_Cap(t *testing.T) {
	var stack *Stack = new(Stack)
	var fibTests = []struct {
		expected int // expected result
	}{
		{0},
	}
	for _, tt := range fibTests {
		actual := stack.Cap()
		if actual != tt.expected {
			t.Errorf("cap() ; actual:%d  ,expected %d", actual, tt.expected)
		}
	}
}
func TestStack_IsEmpty(t *testing.T) {
	var stack *Stack = new(Stack)
	var fibTests = []struct {
		expected bool // expected result
	}{
		{true},
	}
	for _, tt := range fibTests {
		actual := stack.IsEmpty()
		if actual != tt.expected {
			t.Errorf("isEmpty() ; actual:%t  ,expected %t", actual, tt.expected)
		}
	}
}
func TestStack_Pop(t *testing.T) {
	var stack *Stack = new(Stack)
	*stack = append(*stack, "a")
	var fibTests = []struct {
		expected string // expected result
	}{
		{"a"},
	}
	for _, tt := range fibTests {
		actual, err := stack.Pop()
		if actual != tt.expected || err != nil {
			t.Errorf("Pop() ; actual:%v  ,expected %v", actual, tt.expected)
		}
	}
}
func TestStack_Push(t *testing.T) {
	var stack *Stack = new(Stack)
	*stack = append(*stack, "a")
	var fibTests = []struct {
		x string //input
	}{
		{"a"},
	}
	for _, tt := range fibTests {
		stack.Push(tt.x)

	}
}
func TestStack_Top(t *testing.T) {
	var stack *Stack = new(Stack)
	*stack = append(*stack, "a")
	var fibTests = []struct {
		expected string // expected result
	}{
		{"a"},
	}
	for _, tt := range fibTests {
		actual, err := stack.Top()
		if actual != tt.expected || err != nil {
			t.Errorf("Top() ; actual:%v  ,expected %v", actual, tt.expected)
		}
	}
}
