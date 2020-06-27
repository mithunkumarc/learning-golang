import struct from other package  

to export struct , declare with uppercase, and its methods too  

create folder zoo  

          zoo > go mod init "github.com/mithun/zoo"

zoo/animal/tiger.go

        package animal

        import (
          "fmt"
        )

        // upppercase to export
        type Tiger struct {

        }
        
        //uppercase to export
        func (t Tiger) Eat() {
          fmt.Println("tiger eat")
        }
        
zoo/test.go

        package main
        import (
          "github.com/mithun/zoo/animal"
        )
        func main() {
          //package name before struct name : animal.Tiger
          var t animal.Tiger = animal.Tiger{}
          t.Eat();
        }
        
go build  

zoo.exe  
