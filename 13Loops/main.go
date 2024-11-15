package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}
	

    // 1. Type 1

    // for i:=0;i<len(days);i++{
    //     fmt.Println(days[i])
    // }


    // 2: Type 2
    
    // for i := range days{
    //     fmt.Println(days[i])
    // }

    
    // 3: Type 3


    // for index,days := range days{
    //     fmt.Printf("Idex is %v and Value is %v\n",index,days)
    // }
    

    // 4. Type 4


    index :=0

    for index<len(days){
        if days[index] == "Thrusday"{
            break
        }
        index++
        fmt.Println(days[index])

       
    }

}
