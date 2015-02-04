/*
* Change Return Program - The user enters a cost and
* then the amount of money given. The program will figure
* out the change and the number of quarters, dimes, nickels,
* pennies needed for the change.
*/

package main

import (
	"fmt"
)

type money struct {
	quarters,dimes,nickels,pennies int
}

func change(c float64) (money) {
	var amount int = int(c * 100) //change doller to cents
	money := money{0,0,0,0}
	for {
		if amount > 25 {
			money.quarters = amount / 25
			amount = amount % 25
		} else if amount > 10 {
			money.dimes = amount / 10
			amount = amount % 10
		} else if amount > 5 {
			money.nickels = amount / 5
			amount = amount % 5
		} else if amount > 0 {
			money.pennies = amount 
			amount = 0
		} else {
			break
		}
		
	}
	return money
}

func main() {
	cost,given := 0.0, 0.0
	fmt.Println("What's the cost in dollars? ")
	fmt.Scanln(&cost)
	fmt.Println("What's the amount of dollars given? ")
	fmt.Scanln(&given)
	amount := given - cost
	if amount < 0 {
		fmt.Printf("Please ask for $%.2f more from customer.\n", -amount)
	} else {
		fmt.Printf("{quaters, dimes, nickels, pennies}: %v\n", change(amount))
	}
}