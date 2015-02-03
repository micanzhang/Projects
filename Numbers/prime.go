package main

import (
	"fmt"
)

func IsPrime(n int) (bool) {
	if n < 2 {
		return false
	} 
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	n := 0
	fmt.Println("Enter the number to find prime factors of: ")
	fmt.Scanln(&n)
	step := n
	for i:= 2; i<=step; i++ {
		for {
			if step % i != 0 {
				break
			} else {
				if i != n {
					fmt.Printf("%d\t", i)
				}
				step = step / i
			}
		}
	}
	fmt.Println("")
}