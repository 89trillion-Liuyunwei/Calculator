package model

import "testing"

func TestMath_IsNum(t *testing.T) {
	var math Math
	var fibTests = []struct {
		s        string // input
		i        int    //input
		expected bool   // expected result
	}{
		{"a", 0, false},
		{"z", 0, false},
		{"@", 0, false},
		{"！", 0, false},
		{"%", 0, false},
		{"，", 0, false},
		{"$", 0, false},
		{" ", 0, false},
		{"/", 0, false},
		{"*", 0, false},
		{"+", 0, false},
		{"-", 0, false},
		{"^", 0, false},
		{"0", 0, true},
		{"13", 0, true},
		{"322", 0, true},
		{"343", 0, true},
		{"24", 0, true},
		{"6", 0, true},
		{"72", 0, true},
		{"8", 0, true},
		{"9", 0, true},
	}
	for _, tt := range fibTests {
		actual := math.IsNum(tt.s)
		if actual != tt.expected {
			t.Errorf("IsNum(%s) = %t; expected %t", tt.s, actual, tt.expected)
		}
	}
}

func TestMath_Operator(t *testing.T) {
	var math Math
	var fibTests = []struct {
		x        int    // input
		y        int    // input
		operator string //input
		expected string // expected result
	}{
		{1, 2, "+", "3"},
		{100, 22, "-", "78"},
		{100, 2, "*", "200"},
		{100, 2, "/", "50"},
		{100, 2, "%", "0"},
		{100, 2, "^", "10000"},
	}
	for _, tt := range fibTests {
		actual, _ := math.Operator(tt.x, tt.y, tt.operator)
		if actual != tt.expected {
			t.Errorf("Operator(%d,%d,%s) = %s; expected %s", tt.x, tt.y, tt.operator, actual, tt.expected)
		}
	}
}

func TestMath_IsLower(t *testing.T) {
	var math Math
	var fibTests = []struct {
		x string // input
		y string // input

		expected bool // expected result
	}{
		{"+", "-", false},
		{"*", "/", false},
		{"+", "/", true},
		{"-", "/", true},
		{"^", "*", false},
		{"^", "/", false},
		{"^", "-", false},
	}
	for _, tt := range fibTests {
		actual := math.IsLower(tt.x, tt.y)
		if actual != tt.expected {
			t.Errorf("isLower(%s,%s) = %t; expected %t", tt.x, tt.y, actual, tt.expected)
		}
	}
}

func TestMath_JudgmentType(t *testing.T) {
	var math Math
	var fibTests = []struct {
		s        string // input
		i        int    //input
		expected bool   // expected result
	}{
		{"a", 0, false},
		{"z", 0, false},
		{"@", 0, false},
		{"！", 0, false},
		{"%", 0, false},
		{"，", 0, false},
		{"$", 0, false},
		{" ", 0, false},
		{"/", 0, false},
		{"*", 0, false},
		{"+", 0, false},
		{"-", 0, false},
		{"^", 0, false},
		{"0", 0, true},
		{"1", 0, true},
		{"2", 0, true},
		{"3", 0, true},
		{"2", 0, true},
		{"6", 0, true},
		{"7", 0, true},
		{"8", 0, true},
		{"9", 0, true},
	}
	for _, tt := range fibTests {
		actual := math.JudgmentType(tt.s, 0)
		if actual != tt.expected {
			t.Errorf("JudgmentType(%s) = %t; expected %t", tt.s, actual, tt.expected)
		}
	}
}
func TestMath_Add(t *testing.T) {
	var math Math
	var fibTests = []struct {
		x        int // input
		y        int // input
		expected int // expected result
	}{
		{21, 2, 23},
	}
	for _, tt := range fibTests {
		actual := math.add(tt.x, tt.y)
		if actual != tt.expected {
			t.Errorf("add(%d,%d) = %d; expected %d", tt.x, tt.y, actual, tt.expected)
		}
	}
}
func TestMath_Sub(t *testing.T) {
	var math Math
	var fibTests = []struct {
		x        int // input
		y        int // input
		expected int // expected result
	}{
		{2, 2, 0},
	}
	for _, tt := range fibTests {
		actual := math.sub(tt.x, tt.y)
		if actual != tt.expected {
			t.Errorf("sub(%d,%d) = %d; expected %d", tt.x, tt.y, actual, tt.expected)
		}
	}
}
func TestMath_Mul(t *testing.T) {
	var math Math
	var fibTests = []struct {
		x        int // input
		y        int // input
		expected int // expected result
	}{
		{2, 2, 4},
	}
	for _, tt := range fibTests {
		actual := math.mul(tt.x, tt.y)
		if actual != tt.expected {
			t.Errorf("mul(%d,%d) = %d; expected %d", tt.x, tt.y, actual, tt.expected)
		}
	}
}
func TestMath_Division(t *testing.T) {
	var math Math
	var fibTests = []struct {
		x        int // input
		y        int // input
		expected int // expected result
	}{
		{2, 2, 1},
	}
	for _, tt := range fibTests {
		actual := math.division(tt.x, tt.y)
		if actual != tt.expected {
			t.Errorf("division(%d,%d) = %d; expected %d", tt.x, tt.y, actual, tt.expected)
		}
	}
}
func TestMath_Modulo(t *testing.T) {
	var math Math
	var fibTests = []struct {
		x        int // input
		y        int // input
		expected int // expected result
	}{
		{2, 2, 0},
	}
	for _, tt := range fibTests {
		actual := math.modulo(tt.x, tt.y)
		if actual != tt.expected {
			t.Errorf("moudlo(%d,%d) = %d; expected %d", tt.x, tt.y, actual, tt.expected)
		}
	}
}
func TestMath_Pow(t *testing.T) {
	var math Math
	var fibTests = []struct {
		x        int // input
		y        int // input
		expected int // expected result
	}{
		{2, 2, 4},
	}
	for _, tt := range fibTests {
		actual := math.pow(tt.x, tt.y)
		if actual != tt.expected {
			t.Errorf("pow(%d,%d) = %d; expected %d", tt.x, tt.y, actual, tt.expected)
		}
	}
}
