package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := int64(439) // converting to binary
	t := strconv.FormatInt(n, 2) 		
	count := 0   
	for _,w := range strings.Split(t,"0") {
		if count < len(w){
			count = len(w)
		}
	}	
	fmt.Println(count)
}

/*
  convert number to binary
  lets say 11011101111 
  split at 0
  get length of each sub string
  largest is the maximum consecutive sequence of 1s

*/
