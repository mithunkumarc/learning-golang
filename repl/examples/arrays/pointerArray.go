package arrays

import "fmt"

func PointerArrayExample() {
	numbers := [...]int{1,2,3}
	pNumbers := &numbers
	fmt.Println("reading from pointer array: same as normal array")
	for _, number := range pNumbers {
		fmt.Println(number)
	}
	fmt.Println(pNumbers) // &[1,2,3]
	multiplier(&numbers, 5) // multiple each element by 5
	fmt.Println("after multiplier")
	fmt.Println(numbers) // [5, 10 ,15]
}

//pass by reference
func multiplier(numbers *[3]int, multiply int) {
	for index, number := range numbers {
		numbers[index] = number * multiply
	}
}

/*main.go

package main

import (
  "main/arrays"
)

func main() {
 arrays.PointerArrayExample()
}
*/