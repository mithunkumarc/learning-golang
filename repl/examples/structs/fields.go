package structs

import (
  "fmt"
)
// type Vertex struct {}

type BankAccount struct {
  Name string
  Balance float64
  AccountNumber string
  password string // private , unexported
}
func AccessModifierExample() {
  var b BankAccount = BankAccount{
                      Name:"madan",
                      Balance:3232.23,
                      AccountNumber: "abc2323",
                      password: "lskjfdsk"}//close here or else error
  fmt.Println(b)
}


/*main.go

package main

import (
  "main/structs"
  "fmt"
)

func main() {
 structs.AccessModifierExample()
 //password is unknown here not exported
 var ba structs.BankAccount = structs.BankAccount{Name:"ranga",Balance:3434,AccountNumber:"sldjf"}
fmt.Println(ba)
}
*/