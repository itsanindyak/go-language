package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)





	fmt.Println("Welcome to pizza app\nPlease rate our pizza bwt 1 and 5")

	input,_ := reader.ReadString('\n')


	numInput,err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(numInput+1)
	}







	


}