package structs

import (
  "fmt"
)

type Car struct {
  brand string
  cost float64
}
func WithPointers() {
  var c Car = Car {brand: "maruti", cost: 433}
  displayCar(&c)
}

func displayCar(cc *Car) {
  fmt.Println(cc)
  fmt.Println(cc.brand, cc.cost)
}