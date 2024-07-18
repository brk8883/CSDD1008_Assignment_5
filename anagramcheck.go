package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(s string) string {
	slice := strings.Split(s, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}

func IsAnagram(str1, str2 string) bool {

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	if len(str1) != len(str2) {
		return false
	}

	return SortString(str1) == SortString(str2)
}

func main() {
	var str1, str2 string
	fmt.Println("Enter first string:")
	fmt.Scanln(&str1)
	fmt.Println("Enter second string:")
	fmt.Scanln(&str2)

	if IsAnagram(str1, str2) {
		fmt.Println("The strings are anagrams.")
	} else {
		fmt.Println("The strings are not anagrams.")
	}
}
