package main

import "fmt"

func main() {
	fmt.Println("Array in go")

	// declare
	var fruit [4]string

	fruit[0] = "Apple"
	fruit[1] = "Apple"
	fruit[3] = "Apple"
	var vegList = [4]string{"Potato", "Beans", "Mushroom"}

	fmt.Println("List is", fruit)
	fmt.Println("List is", vegList)
}
