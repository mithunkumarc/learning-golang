package enum
import (
  "fmt"
)
/*
no keyword called enum in go
use type keyword to create enum
*/
type day int
/*
enum index by default start with 0
customized using keyword iota 
iota : some greek letter
*/
const (
  /*if your skip iota all values
  gets index 0
  enums/const are by default exported: public
  */
  MONDAY day = 0 + iota // iota increments index
  TUESDAY
  WEDNESDAY
  THURSDAY
  FRIDAY
  SATURDAY
  SUNDAY
)

func EnumExammple() {
  fmt.Println(MONDAY)
  fmt.Println(TUESDAY)
  fmt.Println(WEDNESDAY)
  fmt.Println(THURSDAY)
  fmt.Println(FRIDAY)
  fmt.Println(SATURDAY)
  fmt.Println(SUNDAY)

  fmt.Println(MONDAY == 0) // true
}