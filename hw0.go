package main

import "fmt"

type MyStruct struct {
	myInt int
	myString string
}

func (myStruct MyStruct) getMyInt() int {
	return myStruct.myInt
}

func (myStruct MyStruct) getMyString() string {
	return myStruct.myString
}

func main() {
	fmt.Println("Hello, World!")
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	if(n%3==0 && n%5==0) {
		return "FizzBuzz"
	} else if(n%3 == 0) {
		return "Fizz"
	} else if(n%5 == 0) {
		return "Buzz"
	}
	return ""
}

// // IsPrime checks if the number is prime.
// // You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	if(n < 2) {
		return false
	}

	for i := 2; i < n-1; i++ {
		if(n%i == 0) {
			return false
		}
	}
	return true
}

// // IsPalindrome checks if the string is a palindrome.
// // A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	reverseS := ""

	for i := len(s)-1; i >= 0; i-- {
		reverseS += string(s[i])
	}

	for i := 0; i < len(s); i++ {
		if(reverseS[i] != s[i]) {
			return false
		}
	}
	return true
}