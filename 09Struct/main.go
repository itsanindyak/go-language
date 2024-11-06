package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	fmt.Println("Struct in go")
	anindya := User{"anindya","iankoley@gmail.com",18,false}

	fmt.Println(anindya)
	fmt.Printf("%+v\n",anindya)
	fmt.Printf("Name is %s and Email is %s",anindya.Name,anindya.Email)
}
