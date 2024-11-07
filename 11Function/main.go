package main

import "fmt"

func main() {
	fmt.Println("Hello Golang")
	greeter()
    fmt.Println(add(12,24))
}

func greeter() {
	fmt.Println("Namaste Golang")
}

func add(a int, b int) int {
	return a + b
}
