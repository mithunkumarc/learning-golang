package main

import (
	"fmt"
)

func main() {
	//pointers
	i := 0
	fmt.Println(&i) // prints address
	fmt.Printf("%T\n",i) // int
	fmt.Printf("%T\n",&i) // *int ,pointer to an int
	
	//declare pointer
	var j *int = &i
	fmt.Println("j ",j) //adress
	fmt.Println("value of i ", *j) // 0, get value stored at an address, dereferencing
	fmt.Printf("type of j %T\n",j) // *int
	
	//change value using pointer
	*j = 100
	fmt.Println(" value at i ", i) // 100
	
	
}

