package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("Defer is runing")
	fmt.Println("Defer is not present here")
	fmt.Println("3.. 2.. 1..")

	// Continue and break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue
		if i == 2 {
			fmt.Println("It's 2")
			continue
		}

		// Break
		if i == 9 {
			fmt.Println("Break is happening")
			break
		}

	}
}
