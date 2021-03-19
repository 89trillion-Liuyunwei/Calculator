package model

import (
	"testing"
)

func TestStringently(t *testing.T) {

	var fibTests = []struct {
		s        string // input
		i        int
		expected bool // expected result
	}{
		{"a", 0, false},
		{"z", 0, false},
		{"@", 0, false},
		{"！", 0, false},
		{"%", 0, true},
		{"，", 0, false},
		{"$", 0, false},
		{" ", 0, true},
		{"/", 0, true},
		{"*", 0, true},
		{"+", 0, true},
		{"-", 0, true},
		{"^", 0, true},
		{"0", 0, true},
		{"1", 0, true},
		{"2", 0, true},
		{"3", 0, true},
		{"4", 0, true},
		{"6", 0, true},
		{"7", 0, true},
		{"8", 0, true},
		{"9", 0, true},
	}
	for _, tt := range fibTests {
		actual := Stringently(tt.s, tt.i)
		if actual != tt.expected {
			t.Errorf("Stringently(%s,%d) = %t; expected %t", tt.s, tt.i, actual, tt.expected)
		}
	}

}
