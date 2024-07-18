package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"listen", "silent", true},
		{"triangle", "integral", true},
		{"apple", "pale", false},
		{"Apple", "pplae", true},
		{"rat", "car", false},
		{"hello", "billion", false},
		{"Listen", "Silent", true},
	}

	for _, test := range tests {
		result := IsAnagram(test.str1, test.str2)
		if result != test.expected {
			t.Errorf("IsAnagram(%q, %q) = %v; expected %v", test.str1, test.str2, result, test.expected)
		}
	}
}
