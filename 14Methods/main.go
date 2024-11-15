package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User)getStatus(){ // If any method is just print the data of struct,then not use pointer of struct.
	if(u.Status){
		fmt.Println(u.Name,"is online.")
	}else{
		fmt.Println(u.Name,"is not online.")
	}
}
func (u *User) toggleStatus() {   // If any method is used to change the data of struct then user pointer of struct.
	u.Status = !u.Status
	fmt.Println("Status toogle succesfully")
}

func (u *User)newMail(mail string){
	oldMail := u.Email
	u.Email=mail
	fmt.Println("Email changed",oldMail,"to",u.Email)
}


func main() {
	fmt.Println("Struct in go")
	anindya := User{"anindya koley","iankoley@gmail.com",18,false}

	anindya.getStatus()
	anindya.toggleStatus()
	anindya.newMail("anindya@koley.dev")
}
