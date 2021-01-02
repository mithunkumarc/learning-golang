package controlstatements

import (
  "fmt"
)

func SwitchExample(flag bool) {
    /*no need of break statement for each case*/
    /*in below example: default is unreachable*/
    switch flag {
        case true: fmt.Println("i am truth")
        case false: fmt.Println("i am false")
        default: fmt.Println("i am nobody")
  }
}

/*main.go

package main
import (
 "main/controlstatements"
)

func main() {
 controlstatements.SwitchExample(true) 
 controlstatements.SwitchExample(false) 
}

*/