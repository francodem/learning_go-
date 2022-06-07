package main

import (
	"fmt"
)

func squareArea(side float64) float64 {
	return side * side
}

func sum(a float64, b float64) float64 {
	return a + b
}

func minus(a float64, b float64) float64 {
	return a - b
}

func multiply(a float64, b float64) float64 {
	return a * b
}

func divide(a float64, b float64) float64 {
	return a / b
}

func modulus(a int, b int) int {
	return a % b
}

func increment(value int) int {
	value++
	return value
}

func decrement(value int) int {
	value--
	return value
}

func rectangleArea(side float64, base float64) float64 {
	return side * side
}

func trapezoidArea(B float64, b float64, h float64) float64 {
	return ((B + b) * h) / 2
}

func circleArea(r float64) float64 {
	const pi float64 = 3.1415
	return pi * r * r
}

func main() {
	fmt.Printf("Square Area = %.2f\n", squareArea(4))
	fmt.Printf("Sum result = %.2f\n", sum(5, 10))
	fmt.Printf("Minus result = %.2f\n", minus(5, 10))
	fmt.Printf("Multiply result = %.2f\n", multiply(20, 2.5))
	fmt.Printf("Divide result = %.2f\n", divide(150, 2))
	fmt.Printf("Remainder result = %d\n", modulus(100, 5))
	fmt.Printf("Increment result = %d\n", increment(9))
	fmt.Printf("Decrement result = %d\n", decrement(10))
	fmt.Printf("Rectangle area result = %.2f\n", rectangleArea(5, 10))
	fmt.Printf("Trapezoid area result = %.2f\n", trapezoidArea(8, 5, 2))
	fmt.Printf("Circle area result = %.2f\n", circleArea(8))
}
