package main

import "fmt"

func main() {
	fmt.Println("Good night")

	var ptr1 *int
	fmt.Println("Value of pointer is",ptr1)


	count := 23

	

	var ptr2 = &count
	fmt.Println("Value of pointer is",ptr2)
	fmt.Println("Value of pointer is",*ptr2)




}