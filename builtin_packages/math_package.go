package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	Separator(" pi ")
	// constant
	fmt.Println("pi : ",math.Pi)

	Separator(" Ceil ");	
		
	// Ceil : next nearest integer. 
	fmt.Println("math.Ceil(25.0) : ", math.Ceil(25.0)) // return same number 25
	fmt.Println("math.Ceil(25.1) : ", math.Ceil(25.1)) // from 25.1 to 25.9 , returns 26
	
	Separator(" Dim ")
	
	// Dim : returns difference of x-y if x should be greater than y, else return 0
	fmt.Println("math.Dim(10,6) : ", math.Dim(10,6)) // x > y, x - y = 4
	fmt.Println("math.Dim(-5,10) : ", math.Dim(-5,10)) // x < y, so 0
	fmt.Println("math.Dim(-5, -4) : ",math.Dim(-5, -4)) // x < y, so 0
	fmt.Println("math.Dim(-5, -7) : ", math.Dim(-5, -7)) // x > y, x - (-y) = 2
	
	Separator(" Floor ")
	
	// Floor : its like opposite to Ceil, round off number to present or previous number
	fmt.Println("math.Floor(1.9) : " , math.Floor(1.1)) // roundoff to previous number 1
	fmt.Println("math.Floor(1.9) : " , math.Floor(1.9)) // roundoff to previous number 1
	fmt.Println("math.Floor(12.0) ", math.Floor(12.0)) // roundoff to same number 12
	
	Separator(" pow ")
	
	// pow(x,y) : x ** y, x to the power of y
	fmt.Println("math.Pow(5,2) : ",math.Pow(5,2))

	Separator(" Round ")
	
	// Round : roundoff to nearest integer
	fmt.Println("math.Round(1.1) : ", math.Round(1.1)) // 1
	fmt.Println("math.Round(1.5) : ", math.Round(1.5)) // 2
	fmt.Println("math.Round(1.6) : ", math.Round(1.6)) // 2
	fmt.Println("math.Round(1.9) : ", math.Round(1.9)) // 2

	Separator("sqrt")
	// sqrt : returns square root of number
	fmt.Println("math.Sqrt(100) : ", math.Sqrt(100)) // 10
}

func Separator(title string) {
	pattern := strings.Repeat("*", 10)
	fmt.Println(pattern , title, pattern)
}


/* output 
**********  pi  **********
pi :  3.141592653589793
**********  Ceil  **********
math.Ceil(25.0) :  25
math.Ceil(25.1) :  26
**********  Dim  **********
math.Dim(10,6) :  4
math.Dim(-5,10) :  0
math.Dim(-5, -4) :  0
math.Dim(-5, -7) :  2
**********  Floor  **********
math.Floor(1.9) :  1
math.Floor(1.9) :  1
math.Floor(12.0)  12
**********  pow  **********
math.Pow(5,2) :  25
**********  Round  **********
math.Round(1.1) :  1
math.Round(1.5) :  2
math.Round(1.6) :  2
math.Round(1.9) :  2
********** sqrt **********
math.Sqrt(100) :  10
*/
