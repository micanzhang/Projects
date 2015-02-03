package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Enter number of digits to round PI to: ")
	n := 0
	fmt.Scanln(&n)
	fmt.Printf("%.*f\n", n, 4*(4*math.Atan(1.0/5.0)-math.Atan(1.0/239.0)))

}
