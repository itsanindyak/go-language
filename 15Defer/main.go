package main

import "fmt"

func main() {

	// defer fmt.Println("1")
	// defer fmt.Println("2")
	// defer fmt.Println("3")
	// defer fmt.Println("4")
	// fmt.Println("5")
	// fmt.Println("6")
	// defer fmt.Println("7")
	// defer fmt.Println("8")
	// defer fmt.Println("9")
	//  result are :
	//  5 6 9 8 7 4 3 2 1

	fmt.Println(complexExample())
	fdefer()

}

func complexExample() (value int) {
	defer func() {
		value += 1 // Third: value becomes 3
		fmt.Println(value,"1")
	}()
	defer func() {
		value *= 2 // Second: value becomes 2
		fmt.Println(value,"2")
	}()
	return 1 // First: value starts as 1
}
func fdefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
