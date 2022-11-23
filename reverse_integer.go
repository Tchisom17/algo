package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-467))
}

func reverse(x int) int {
	newNum := 0
	sign := 1

	if x < 0 {
		sign = -1
		x = -x
	}

	for x > 0 {
		remainder := x % 10
		newNum = newNum*10 + remainder
		x /= 10
	}

	if (newNum*sign) > math.MaxInt32 || (newNum*sign) < math.MinInt32 {
		return 0
	}

	return newNum * sign
}
