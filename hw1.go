// Homework 1: Finger Exercises
package main

import (
	// "fmt"
	"sort"
	"strings"
)

// func main() {
// }

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	phone1 := strings.Replace(phone, "-", "", -1)
	phone2 := strings.Replace(phone1, " ", "", -1)

	retPhone := "(" + string(phone2[0:3]) + ") " + string(phone2[3:6]) + "-" + string(phone2[6:10])
	return retPhone
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	s1 = strings.ReplaceAll(s1, " ", "")
	s2 = strings.ReplaceAll(s2, " ", "")

	if len(s1) == len(s2) {
		sl1 := strings.Split(s1, "")
		sl2 := strings.Split(s2, "")
	
		sort.Strings(sl1)
		sort.Strings(sl2)

		for i := 0; i < len(s1); i++ {
			if sl1[i] != sl2[i] {
				return false
			}
		}
		return true
	}

	return false
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	var evens []int

	for i := 0; i < cap(e); i++ {
		if(e[i]%2 == 0) {
			evens = append(evens, e[i])
		}
	}

	return evens
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	sum := 0

	for i := 0; i < cap(e); i++ {
		sum += e[i]	
	}

	return sum
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	uniques := []int{}

	for i := 0; i < cap(e); i++ {
		count := 0
		for j := 0; j < cap(e); j++ {
			if(e[i] == e[j]) {
				count++
			}
		}
		if(count <= 1) {
			uniques = append(uniques, e[i])
		}
	}
	return uniques
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	m := make(map[int]string)

	for k, v := range kv {
		m[v] = k
	}

	return m
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	//Doesn't work with UTF-8
	m := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		item := []rune(s)[i]
		a := m[item]
		if a == 0 {
			m[item] = 1
		} else {
			m[item]++
		}
	}

	keys := make([]rune, len(m))

	for i := range m {
		keys = append(keys, i)
	}

	for i := 0; i < len(keys); i++ {
		if m[keys[i]] < k {
			delete(m, keys[i])
		}
	}

	return m
}
