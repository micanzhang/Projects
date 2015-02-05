/*
* Fibonacci Sequence - Enter a number and have the
* program generate the Fibonacci sequence to that number
* or to the Nth number
 */

package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Println("How many numbers do you need? ")
	n := 0
	fmt.Scanln(&n)
	numbers := []int{}
	for i := 0; i < n; i++ {
		numbers = append(numbers, fibonacci(i))
	}
	fmt.Println(numbers)
}
