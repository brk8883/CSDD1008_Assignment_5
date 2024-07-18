package main

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	// Define test cases
	tests := []struct {
		input    float64
		expected float64
		hasError bool
	}{
		{4, 2, false},
		{9, 3, false},
		{16, 4, false},
		{0, 0, false},
		{-1, 0, true},
		{7, 2.64575131106, false},
	}

	for _, test := range tests {
		result, err := Sqrt(test.input)

		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for input %v, but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Did not expect error for input %v, but got %v", test.input, err)
			}
			// Check if the result is approximately equal to the expected value
			if math.Abs(result-test.expected) > 1e-9 {
				t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
			}
		}
	}
}
