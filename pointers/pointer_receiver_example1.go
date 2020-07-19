#### method accepts obj as pointer : Pointer receivers

        package main

        import (
          "fmt"
          "math"
        )

        type Vertex struct {
          X, Y float64
        }

        func (v Vertex) Abs() float64 {
          return math.Sqrt(v.X*v.X + v.Y*v.Y)
        }

        func (v *Vertex) Scale(f float64) {   //if u skip * in this line, output will be 5
          v.X = v.X * f
          v.Y = v.Y * f
        }
        
        func main() {
          v := Vertex{3, 4}
          // if u skip * in line 20 func(v Vertex), object is treated as value, so it is not reflected here, output 5
          v.Scale(10) 
          fmt.Println(v.Abs()) //50
        }
        
        // The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
