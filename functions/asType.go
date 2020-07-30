package main

import (
	"fmt"
)

func world(t string) {
	fmt.Println(t)
}

func main() {
	var hello func(string) = world;
	fmt.Printf("type of hello is %T \n",hello) // func(string)
	hello("golang")
}
