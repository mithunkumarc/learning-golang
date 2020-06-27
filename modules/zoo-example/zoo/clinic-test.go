package main
import (
	"fmt"
	"github.com/mithun/zoo/animals"
	"github.com/mithun/zoo/birds"
	)
func main() {
	fmt.Println("hello world")
	animals.LionEat()
	animals.TigerEat()
	birds.EmuEat()
	birds.ParrotEat()
}