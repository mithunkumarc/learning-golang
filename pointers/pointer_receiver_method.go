package main

import (
	"fmt"
)

type Employee struct {
	name string
}

// no need to send reference example: (&e).changeName()
func (e *Employee)changeName() {
	e.name = "anjan"
	fmt.Println(e);
}

func main() {
	e := Employee{name: "rajat"}
	e.changeName() //by default reference is sent
	fmt.Println(e.name)
}

/*
&{anjan}
anjan

*/
