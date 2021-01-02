package arrays
import "fmt"

// index/position can be assined to array elements
func WithIndexExample() {

	words := [...]string{1:"hello",3:"world"}
	fmt.Println(len(words)) // 4
	fmt.Println(words) // ["" "hello" "" "world"]
	fmt.Println("words[0] ",words[0]) // empty
	fmt.Println("words[1] ",words[1]) // "hello"
	fmt.Println("words[2] ",words[2]) // emtpy
	fmt.Println("words[3] ",words[3]) // "world"
}

/*main 

package main

import (
  "main/arrays"
)

func main() {
 arrays.WithIndexExample()
}
*/