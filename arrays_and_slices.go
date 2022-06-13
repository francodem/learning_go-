package main

import "fmt"

func main() {
	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	// Print all array
	fmt.Println(array)
	// Print only one item of the array
	fmt.Println(array[1])

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Slice methods
	// Return an element
	fmt.Printf("A) %d\n", slice[0])
	// Return elements from point a to point B (point B is the limit)
	fmt.Printf("B) %d\n", slice[:3])
	fmt.Printf("C) %d\n", slice[2:4])
	fmt.Printf("D) %d\n", slice[4:])

	// Append into a slice
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append new list
	new_slice := []int{8, 9, 10}
	slice = append(slice, new_slice...)
	fmt.Println(slice)
}
