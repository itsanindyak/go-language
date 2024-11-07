package main

import (
	"fmt"
	"math/rand"
)

func main() {
	
	diceNumber := rand.Intn(6)+1

	fmt.Println("Valu of dice is ----> ",diceNumber)

	switch diceNumber{
	case 1:
		fmt.Println("You can open")
	case 2:
		fmt.Println("You can mode 2 spot")
	case 3:
		fmt.Println("You can mode 3 spot")
		fallthrough
	case 4:
		fmt.Println("You can mode 4 spot")
	case 5:
		fmt.Println("You can mode 5 spot")
		fallthrough        // fallthrough ---> it run the next case also
	case 6:
		fmt.Println("You can mode 6 spot and roll dice again.")
	default:
		fmt.Println("Invalid")
	}
}
