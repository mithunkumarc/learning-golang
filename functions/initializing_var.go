package main

import (
	"fmt"
)

func main() {
	//initializing variable using anonymous function
	x := func () int {
		return 23
	}()
	fmt.Println(x) //23
}
