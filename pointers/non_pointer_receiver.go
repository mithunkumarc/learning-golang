package main

import (
	"fmt"
)

type Phone struct {
	number int
}

//any struct which as call() attached to it becomes Device interface
type Device interface {
	call()
}
//phone become device
func (p Phone) call() {
	fmt.Println("calling number ", p.number) 
}

// non pointer receiver receives , non pointer value as well as pointer value
// phone instance can be sent to Dialer
func Dialer(d Device) {
	d.call()
}

func main() {
	p := Phone{number : 123}
	Dialer(p) // non pointer value ,works
	Dialer(&p) // pointer value, works
	fmt.Printf("%T\n",&p) // pointer *main.Phone ,  main : package name
}

//output : 
//calling number  123
//calling number  123
//*main.Phone
