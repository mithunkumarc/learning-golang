package main

import (
	"fmt"
)

func main() {
	//callback function : sending function as argument
	x := 5
	//arguments function and element
	y := square_calculator(square, x)
	fmt.Println(y)
}

//this function is being sent as argument to square_calculator
func square(y int) int {
	return y * y
}

//first arugment is function
func square_calculator(f func(x int) int , i int) int {
	return f(i)
}

// 25
