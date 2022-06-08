package main

import "fmt"

func main() {
	module := 5 % 2
	switch module {
	case 0:
		fmt.Println("It's even")
	default:
		fmt.Println("It's odd")
	}

	value := 200
	switch {
	case value > 100:
		fmt.Println("It's greater than 100")
	case value < 0:
		fmt.Println("It's less than 0")
	default:
		fmt.Println("Without condition")
	}

}
