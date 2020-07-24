package main

import (
	"fmt"
)

func greet(name string) string {
	return "hello, " + name
}

func main() {
	messanger := greet
	fmt.Println(messanger("ranju"))
	fmt.Printf("%T \n ", messanger)// func(string) string
}
