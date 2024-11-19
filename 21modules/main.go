package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeter("anindya")


	r:= mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")



	log.Fatal(http.ListenAndServe(":3400",r))

}

func greeter(name string) {
	fmt.Println("Hey", name, "Welcome to mod")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hey Anindya</h1>"))
}
