package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name :")
	input,_ :=reader.ReadString('\n')
	input = strings.TrimSpace(input)
	greeter(input)
	add,isAdd:= proAdd(1,2,3,4,5)

	if(isAdd){
		fmt.Println(add)
	}else{
		fmt.Println("Problem to get answer from the function")
	}
}

func greeter(name string) {
	fmt.Println("Namaste", name, ", Lets learn function in goloang")
}

func add(a int, b int) int {
	return a + b
}

func proAdd(values ...int) (int,bool) {
	// values is now a slices
	total := 0
	for i := range values {
		total = add(total, values[i])
	}

	return total,false
}
