package main

import "fmt"

type salesPerWeek[T string | int] struct {
	sales []T
}

func (s salesPerWeek[T]) printSales() { // this function only recived integer slice or string slice
	for i, item := range s.sales {
		fmt.Println(i+1, "Day -->", item)
	}
}

// func printSlice[T int | string](s []T) { // this function only recived integer slice or string slice
// 	for _, item := range s {
// 		fmt.Println(item)
// 	}
// }

func main() {

	mySale := salesPerWeek[string]{
		sales: []string{"A", "S", "D", "F", "G", "H"},
	}

	mySale.printSales()
	// var score = []int{12, 23, 34, 45, 67, 78}
	// var stringScore = []string{"A", "S", "D", "F", "G", "H"}
	// printSlice(score)

}
