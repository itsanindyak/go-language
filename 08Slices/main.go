package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Go Slices")

	var fruitList = []string{"Apple","Pinapple","Orange"}
	// fmt.Printf("%T\n", fruitList)

	fruitList = append(fruitList, "mango", "Banana")
	fruitList= fruitList[1:2]

	// fmt.Println(fruitList)




	//	 use make --> make allocated memory

	highScore := make([]int,4)

	highScore[0]=127
	highScore[1]=1
	highScore[2]=1
	highScore[3]=1

	highScore = append(highScore, 555,666,777)

	sort.Ints(highScore)

	//fmt.Println(highScore)


	// how to remove a value from slices based on index

	var coursess = []string{"Javascript","C++","C","Python","Go","Typescript"}
	index := 2
	coursess = append(coursess[:index],coursess[index+1:]... ) //  "  ...  " ---> used to pass extra argument in a function


	fmt.Println(coursess)


















}
