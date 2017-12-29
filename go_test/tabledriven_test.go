package main

import "testing"

func plus(x, y int) int {
	return x + y
}

func TestPlus(t *testing.T) {
	var tests = []struct {
		x      int
		y      int
		expect int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{3, 2, 5},
	}

	for _, tt := range tests {
		actual := plus(tt.x, tt.y)
		if actual != tt.expect {
			t.Errorf("add(%d, %d): exprect %d, actual %d", tt.x, tt.y, tt.expect, actual)
		}
	}
}
