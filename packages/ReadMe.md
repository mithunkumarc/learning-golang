Read 
 
       golangbot.com/go-packages/


#### How to import packages

        creating project : go mod init calculator
        
        user defined package mathematics
        mathematics // folder
          adder.go //package mathematics
        
        main.go // entry file

        /* code in main.go
        import (
	         m "calculator/mathematics"
        	 "fmt"
         )
        */
      

### folder structure , you can design folder structure like java. but for golang all the go files inside the a package merge into single big file(not literally). 
### Even if you don't create folder for package , it is enough that each source file declare package name in the beginning, it still works.
### no need to declare import statement if one file importing from another file from same package


    go mod init github.com/zoo

    package name : animal
                      tiger.go
                          eat()
                      lion.go
                          run()
                          
                   above eat() and run() can be called using package name
                   
                   in main.go , import "github.com/zoo/animal"
                   eg: animal.eat()
                       animal.run()
                 
                 * look like files lose identity , their identity is just package name
                 
                 
                 
### to export functions and variables start with upper case letter ,so that other files can import and use, doest matter importing file belong to same or different package


                      animal
                         tiger.go
                         clicic.go
                         
                     tiger and clinic both in same package, but still to export function ,declare with upper case letter    
                         
                    
#### clinic.go


                  package animal
                  import "fmt"
                  func CheckUp() {
                    fmt.Println("animal clinic test")
                  }
                  
#### tiger.go
                    
                    package animal
                    // export func using name starting letter with upper case
                    func Eat() string{
                      return "eating"
                    }
                    func GoClinic() {
                      CheckUp()
                    }

---


#### creating folder for package name is optional but its good to create to see project structure and grouping source files


        ex :  good to create folder name called animal

                      tiger.go 
                          package animal
                          // ...


                      clinic.go
                          package animal
                          // ...    

