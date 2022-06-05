package main

import "fmt"

func main() {
	// Vars declaration
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf
	name := "Franco"
	learned := 2
	fmt.Printf("%s has more than %d hours learning about Go today\n", name, learned)
	// If we don't know about the data type
	fmt.Printf("%v has more than %v hours learning about Go today\n", name, learned)

	// Sprintf
	message := fmt.Sprintf("%s has more than %d hours learning about Go today", name, learned)
	fmt.Println(message)

	// Data type
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("learned: %T", learned)
}
