package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}
func isPalindrome(x int) bool {
	newNum := 0
	oldNum := x

	if x < 0 {
		return false
	}

	for x > 0 {
		remainder := x % 10
		newNum = newNum*10 + remainder
		x /= 10
	}

	if newNum == oldNum {
		return true
	}
	return false
}
