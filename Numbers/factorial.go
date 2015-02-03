package main

import (
	"fmt"
)

func FactLoop(n int) int {
	fact := 1
	for ;n > 0; n-- {
		fact *= n
	}
	return fact
}

func FactRecursion(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * FactRecursion(n-1)
	}
}

func main() {
	fmt.Println("Enter a positive number: ")
	n := 0
	fmt.Scanln(&n)
	if n < 0 {
		fmt.Println("Invalid number")
	} else {
		fmt.Printf("Factorial of %d by loops is %d\n", n, FactLoop(n))
		fmt.Printf("Factorial of %d by recursion is %d\n", n, FactRecursion(n))
	}
}
