package main

import "fmt"

func main() {
	var username string ="Anindya"
	fmt.Println(username)
	fmt.Printf("Type of user name %T \n",username)

	var isGood bool = true
	fmt.Println(isGood)
	fmt.Printf("Type of isGood %T \n",isGood)

	var name float64
	fmt.Println(name)
	fmt.Printf("type = %T ",name)


	// implicit declare
	
	var numberOfApi = 12.0
	fmt.Println(numberOfApi)
	fmt.Printf("Type of isGood %T \n",numberOfApi)

	// default value 
	var numberOfApp string
	fmt.Println(numberOfApp)
	fmt.Printf("Type of isGood %T \n",numberOfApp)


	// no var declare 

	numberOfRequest := "koley"
	fmt.Println(numberOfRequest)
	fmt.Printf("Type of isGood %T \n",numberOfRequest)

	// const 

	const loginToken = "ghhheriiksnkdfjduiop"
	fmt.Println(loginToken)
	
}