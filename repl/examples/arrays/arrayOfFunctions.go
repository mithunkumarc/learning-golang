package arrays

import "fmt"

func ArrayOfFunctions() {
	
	var t1 = func(){fmt.Println("function t1")}
	var t2 = func(){fmt.Println("function t2")}
	
	functions := [2]func() {t1, t2} // function type is : func()
  // if you don't care index, use _
  // unused var throws error in golang
	for _ , function := range functions {
		function()
	}
	fmt.Println(functions) // addresses  [0x491740 0x491750]
}

