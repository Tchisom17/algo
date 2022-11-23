package main

import "fmt"

func main() {
	fmt.Println(reverse(-3467))
}

func reverse(x int) int {
	newNum := 0
	sign := 1

	if x < 0 {
		sign = -1
	}

	if x < 0 {
		x = -x
	} else {
		x = x
	}

	for x > 0 {
		remainder := x % 10
		newNum = newNum*10 + remainder
		x /= 10
	}
	if (newNum*sign) > 2147483647 || (newNum*sign) < -2147483648 {
		return 0
	}

	return newNum * sign
}
