package main

import (
	"fmt"
)

type Phone struct {
	number int
}

//receiving object as pointer argument
func (p *Phone) changeNumber() {
	p.number = 234
	fmt.Println("number change inside changeNumber function ", p.number) // 234
}

func main() {
	p := Phone{number : 133}
	p.changeNumber()
	fmt.Println("in main, number of phone is changed number ",p.number) // 234
}
