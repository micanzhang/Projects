//Next Prime Number - Have the program find prime
//numbers until the user chooses to stop asking for
//the next one.

package main

import (
	"fmt"
	"strings"
)

func NextPrime(current int) int {
	i := 2
	current += 1
	for {
		if i >= current {
			break
		} else if current%i == 0 {
			current += 1
		} else {
			i += 1
		}
	}
	return current
}

func main() {
	currentPrime := 2
	for {
		response := ""
		fmt.Println("Do you want the next prime? (Y/N) ")
		fmt.Scanln(&response)
		if strings.ToLower(response) == "n" {
			break
		} else {
			fmt.Println(currentPrime)
			currentPrime = NextPrime(currentPrime)
		}
	}
}
