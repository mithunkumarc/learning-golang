package main

import (
	"fmt"
)

func Swap(x int , y int) (int, int) {
	return y, x //returning multiple values
}

func getThreeIncrement(r int) (int, int, int) {	
	return r + 1, r + 2, r + 3
}

func main() {
	x, y := 4, 6
	x, y = Swap(x, y)
	fmt.Println("x : ", x, "y : ", y)
  	a, b, c := getThreeIncrement(1)
	fmt.Println(a, b, c)
}

/*

x :  6 y :  4
2 3 4


*/
