package main

import (
	"fmt"
)
// supports index based

func main() {
	for i := range 10 {
		fmt.Println(i)
	}

	for range 10 {
		fmt.Println("do something") // prints 10 times
	}

}
