package main

import (
	"fmt"
)

func main() {
	a := 100
	b := 200
	swap(&a,&b)
	fmt.Println("a : ",a, " b : ", b)	
}

func swap(x , y *int) {
	z := *x // dereferencing, value stored at mem loc pointed by x assigned to z
	// dont use := again here , its only used at initializing for first time
	*x = *y // value stored at mem loc pointed by y is pushed to mem loc pointed by x
	*y = z // value stored in z pushed to mem loc pointed y y
}
