package main

import (
	"fmt"
)

type rectangle struct {
	height, width int;
}

func main() {	
	var t rectangle = rectangle{height:10, width:10}; // key value
	var p rectangle = rectangle{20, 12};  // value : follow order
	fmt.Println(t.height);  // 10
	fmt.Println(p.height);  // 20
}


