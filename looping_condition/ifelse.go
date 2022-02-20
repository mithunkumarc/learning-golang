package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("yes") // yes
	}
	if false {				// false
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
