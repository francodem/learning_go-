package main

import "fmt"

func normalFunction(message string) {
	fmt.Printf("%s on the normal function\n", message)
}

func tripleArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func returnDoubleValues(a int) (c int, d int) {
	return a, a * 2
}

func main() {
	var message string = "This is the message"
	normalFunction(message)
	tripleArgument(1, 5, "hello")

	result := returnValue(4)
	fmt.Println(result)

	result_c, result_d := returnDoubleValues(5)
	fmt.Println(result_c, result_d)

	result_c, _ = returnDoubleValues(8)
	fmt.Println(result_c)
}
