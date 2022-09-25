package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Printf("Enter 1st Number: ")
	fmt.Scan(&x)
	fmt.Printf("Enter 2nd Number: ")
	fmt.Scan(&y)
	fmt.Printf("Sum: %d", x+y)
}
