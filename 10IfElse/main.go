package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("If Else in Go")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the value of a = ")
	input, _ := reader.ReadString('\n')
	a, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 0)

	if a < 10 {
		fmt.Println("Less than 10")
	} else if a > 10 {
		fmt.Println("Greater than 10")
	} else {
		fmt.Println("Equal to 10")
	}



	// special syntax

	if num:=3;num >3{
		fmt.Println("Greater than 3")
	}else if num == 3{
		fmt.Println("Equal to 3")
	}else{
		fmt.Println("Less than 3")
	}
}
