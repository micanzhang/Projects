/* 
* Happy Numbers - A happy number is defined by the
* following process. Starting with any positive integer,
* replace the number by the sum of the squares of its
* digits, and repeat the process until the number equals
* 1 (where it will stay), or it loops endlessly in a
* cycle which does not include 1. Those numbers for which
* this process ends in 1 are happy numbers, while those
* that do not end in 1 are unhappy numbers. Take an input
* number from user, and find first 8 happy numbers from
* that input.
*/

package main

import (
	"fmt"
	"math"
)

func IsHappyNumber(n int) (bool) {
	seen := []float64{}
	for {
		sum_digits := 0.0
		for {
			if n > 0 {
				sum_digits += math.Pow((float64)(n % 10), 2)
				n = n / 10
			} else {
				break
			}
		}
		for _, v := range seen {
			if sum_digits == v {
				return false
			}
		}
		if 1 == sum_digits {
			return true
		} else {
			n = int(sum_digits)
			seen = append(seen, sum_digits)
		}
	}
}

func main() {
	fmt.Println("Start at: ")
	n := 0
	fmt.Scanln(&n)
	happyNumbers := []int{}
	for {
		if len(happyNumbers) < 8 {
			if IsHappyNumber(n) {
				happyNumbers = append(happyNumbers, n)
			}
			n += 1
		} else {
			break
		}
	}
	fmt.Println(happyNumbers)
}