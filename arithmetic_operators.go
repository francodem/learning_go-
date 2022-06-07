package main

import "fmt"

func main() {
	// Square area
	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("Square area: ", squareArea)

	x := 10
	y := 50

	// Sum
	result := x + y
	fmt.Println("Sum: ", result)

	// Minus
	// After the first declaration of result, after that is not necessary to declare := again
	result = x - y
	fmt.Println("Minus: ", result)

	// Multiply
	result = x * y
	fmt.Println("Multiply: ", result)

	// Divide
	result = y / x
	fmt.Println("Divide: ", result)

	// Modulus
	result = y % x
	// 50 % 10 = 5 and is not there a remainder
	fmt.Println("Modulus: ", result)

	// Increment
	x++
	fmt.Println("Increment: ", x)

	// Decrement
	x--
	fmt.Println("Decrement: ", x)

	// Rectangle area
	area := x * y
	fmt.Println("Rectangle area: ", area)

	// Trapezoid area
	area = ((x + y) * 2) / 2
	fmt.Println("Trapezoid area: ", area)

	// Circle area
	const pi float64 = 3.1415
	var radius float64 = 10
	circle_area := pi * (radius * radius)
	fmt.Println("Circle area: ", circle_area)

	// If we want to make a multiplication of a float and a integer, we need to convert the integer to float first
	// and then multiply it with the float variable that we have declared before.
	// After that we declare a new variable of the area.
}
