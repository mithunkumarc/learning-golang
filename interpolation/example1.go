package main

import (
	"fmt"
)

type Employee struct {
	name string
	empId uint32
}

func (e Employee)login() {
	d := fmt.Sprintf("Employee %s with emp id %d loggedin", e.name, e.empId);
	fmt.Println(d);
}

func main() {
	e := Employee{name: "rajat", empId: 23}
	e.login()
}

/*
output : 
Employee rajat with emp id 23 loggedin
*/
