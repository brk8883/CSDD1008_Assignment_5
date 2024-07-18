package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"listen", "silent", true},     // Basic anagram
		{"triangle", "integral", true}, // Anagram with same length
		{"apple", "pale", false},       // Different lengths
		{"Apple", "pplae", true},       // Case insensitivity
		{"rat", "car", false},          // Different characters
		{"hello", "billion", false},    // Different lengths
		{"Listen", "Silent", true},     // Case insensitivity
	}

	for _, test := range tests {
		result := IsAnagram(test.str1, test.str2)
		if result != test.expected {
			t.Errorf("IsAnagram(%q, %q) = %v; expected %v", test.str1, test.str2, result, test.expected)
		}
	}
}
