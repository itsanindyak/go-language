package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to nput"
	fmt.Println(welcome)


	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the price of this product : ")
	input, _ := reader.ReadString('\n')  //  reads the input until it encounters a newline character (\n), including it in the returned string.


	fmt.Println("Your total price is : ",input)
	fmt.Printf("Type of price %T \n",input)
}