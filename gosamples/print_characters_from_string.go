package main

import (
	"fmt"
)

/*
  fmt.Println() prints ascii values
  fmt.Printf() formats ascii value to letter
*/
func main() {
	word := "hello world"
	for i, value := range word {
		fmt.Printf("the index and value is %d and %c \n", word[i], value)
	}
}
