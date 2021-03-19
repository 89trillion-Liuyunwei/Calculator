package model

import (
	"testing"
)

func TestCalculation(t *testing.T) {
	var fibTests = []struct {
		s        string // input
		expected int    // expected result
	}{
		{"(1+2)/1-2^2", -1},
		{"(111-1)/2-2^2", 51},
		{"3^2-1+(2+1)*2", 14},
		{"(12+1^2/0-3O)/2", 0},
		{"(12+1//^2/0-3O)/2", 0},
		{"(12+1,,,kjyu^2/0-3O)/2", 0},
		{"(12+1^^$%2/0-3O)/2", 0},
		{"(12+aaaaaa-3O)/2", 0},
	}
	for _, tt := range fibTests {
		actual, _ := Calculation(tt.s)
		if actual != tt.expected {
			t.Errorf("Calculation(%s) = %d; expected %d", tt.s, actual, tt.expected)
		}
	}

}
func TestPostfixCuration(t *testing.T) {
	s := []string{"5", "10", "2", "1", "/", "2", "^", "*", "-"}
	var fibTests = []struct {
		s        []string // input
		expected int      // expected result
	}{
		{s, -35},
	}
	for _, tt := range fibTests {
		actual, _ := postfixCuration(tt.s)
		if actual != tt.expected {
			t.Errorf("postfixCuration(%v) = %d; expected %d", tt.s, actual, tt.expected)
		}
	}
}
