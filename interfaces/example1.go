// any struct , which defines func declared in interface  
// said to be struct implemented interfaces
package main

import (
    "fmt"
)
type rectangle struct {
	height, width int;
}
// any struct which has area() method attached to it becomes geometry
// imp : function name, parameters and return type
type geometry interface {
	area() int
}

//area() method attached to rectangle
func (r rectangle) area() int{
	return r.height * r.width
}

// test polymorphism, send any struct which has func defined in geometry
func CalArea(g geometry){
	fmt.Println(g.area())
}

func main() {
	r := rectangle{12,45}
	CalArea(r)
	fmt.Println("wait")
}

//source : https://gobyexample.com/interfaces
