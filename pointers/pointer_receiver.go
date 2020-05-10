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
//phone become device, pointer receiver ,important
func (p *Phone) call() {
	fmt.Println("calling number ", p.number) 
}

// pointer receiver receives , only pointer value
// phone instance can be sent to Dialer
func Dialer(d Device) {
	d.call()
}

func main() {
	p := Phone{number : 123}
	//Dialer(p) // non pointer value ,fails, *imp
	Dialer(&p) // pointer value, works
	fmt.Printf("%T\n",&p) // pointer *main.Phone ,  main : package name
}
