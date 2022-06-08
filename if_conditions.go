package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 2
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("It's 1")
	} else {
		fmt.Println("It's not 1")
	}

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Is True")
	}

	if valor1 == 0 || valor2 == 2 {
		fmt.Println("It's True, OR")
	}

	value, err := strconv.Atoi("52")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)
}
