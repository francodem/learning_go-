package main

import "fmt"

func main() {
	// Constants declaration
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println(pi, pi2)
	fmt.Println("pi: ", pi)
	fmt.Println("pi2:", pi2)

	// Integer variables declaration
	// (:=) is the short assignment operator for declare a variable by first time
	base := 12
	var height int = 14
	var area int

	fmt.Println("base:", base, "height:", height, "area:", area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("Zero Values: ", a, b, c, d)

	// Square area
	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("Square area: ", squareArea)
}
