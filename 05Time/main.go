package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	
	fmt.Println(presentTime.Format("2006-01-02 15:04:05"))

	
}