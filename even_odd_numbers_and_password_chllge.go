package main

import "fmt"

func even_check(number int) bool {
	if number%2 == 0 {
		return true
	} else {
		return false
	}
}

func password_check(pass string) bool {
	real_password := "helloworld"
	if pass == real_password {
		return true
	} else {
		return false
	}
}

func main() {
	var number int = 2
	if even_check(number) {
		fmt.Println("It's a even number")
	} else {
		fmt.Println("It's a odd number")
	}

	var pass string = "hellomoon"
	if password_check(pass) {
		fmt.Println("Login correct!")
	} else {
		fmt.Println("Login incorrect!")
	}
}
