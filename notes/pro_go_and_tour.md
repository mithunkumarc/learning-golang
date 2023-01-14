1. variable declaration(and type inference :=)  
2. multiple variable declaration in one line, var a, b int = 5, 6  
3. parsing number in string format to int  
4. how to print character
5. converting chart to string : char to string : string(c)
6. delete elements from slice(no builtin use append)
7. using initialization statement with if and switch 
8. make function : golang builtin package
9. range  : for index, data := range <iterable> 
10. for loop enumerating sequences, string's characters using range 
11. switch : multiple matching values
	
        var numbers string = "112233"
          for _, digit := range numbers {
            switch digit {
            case '1', '2':

12. switch has default break statement for each case 
13. Falling Through to the Next case Statement(missing break function in other languages)
14. default switch value is true 
	
        switch {} same as switch true {}
	
15. comparing arrays : same as python
16. creating slice using make , make([]int, 0, 5) // len(b)=0, cap(b)=5

      https://gobyexample.com/slices
17. copy function

18. creating slice using array : same as python, slice is dynamic in size like list in java 
19. appending slice
20. pointers : dereferencing *p mandatory for builtin 
21. pointers: dereferencing (*person).name optional for user defined(struct)
	
        p := &v
        (*p).x same as p.x
	
22. struct literals(fields)/instance fields : default state(value) same as java like int 0, boolean false
23. struct declaration, The zero value(value of uninitialized) of a slice is nil.
24. slice length and capacity
25. slices of structs(with pointers reference []*Person{{}, {}})
26. function return type , named return, declaring argument list
27. map : create update delete retrive
28. search value in map using key , elem, ok := m[key] , ok : true or false
29. closure
30. struct methods with and without pointer receivers 
31. no set in golang use map, not orders, store keys and sort and iterate map based on keys
32. check diff func speakFunction(p Person) {} vs func (p Person) speak() {}
33. pointer receivers in methods
34. enumerating map, its not in order so collect keys, sort it and read by keys 
35. for struct method use direct reference variable (refvar.methodName())
	
          (method can accept reference/pointer type or value type )
          func (p *Person) speak() or func (p *Person) speak() 

36. if isolated function receives pointer as argument
  
	        someFunction(p *Person) : call this like someFunction(&person_Ref_var)
	

37. In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.????

38. A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

39. Empty interface : An empty interface may hold values of any type. (Every type implements at least zero methods.)
40. Empty interfaces are used by code that handles values of unknown type. 
  
          For example, fmt.Print takes any number of arguments of type interface{}.

41. A type assertion provides access to an interface value's underlying concrete value.

          format : t := i.(T)

          var i interface{} = "hello"
          s := i.(string)
          s, ok := i.(string)
          fmt.Println(s, ok) // hello , true


42. stringer interface : similar to toString() 
	
            type Stringer interface {
              String() string
            }

            func (p Person) String() string {
              return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
            }
	
43. slices of structs two ways

  
            1.storing struct reference/memory location in slice
            products := []*Product{
                {name: "rajat", age: 23},
                {name: "jagat", age: 32},
              }
            [0xc000010030 0xc000010048]
            2.storing struct as value in slice
            products := []Product{
                {name: "rajat", age: 23},
                {name: "jagat", age: 32},
              }
            [{rajat 23} {jagat 32}]

44. error interface  : type error struct { Error() string }

45. panics :

  
          https://www.digitalocean.com/community/tutorials/handling-panics-in-go


46. comparing structs 

          // comparing between value and references(which mem loc pointing to)
          var p1 Product = Product{name: "rajat"}
          var p2 Product = Product{name: "rajat"}
          var p3 *Product = &Product{name: "vinay"}
          var p4 *Product = &Product{name: "vinay"}
          fmt.Println(p1 == p2)   // comparing value : true
          fmt.Println(*p3 == *p4) // comparing value : true
          fmt.Println(p3 == p4)   // comparing reference : false


46. use constructor like function or setter getter for initializing 

47. unexported filed of struct(defined in separate file/module/package)

48. exporting from package/module

            type Product struct { //Product starts with Uppercase letter: exported
              Name, Category string  // Name Category : exported 
              price float64	// price unexported field
            }

            // constructor
            func NewProduct(name, category string, price float64) *Product {
              return &Product{ name, category, price }
            }

            func (p *Product) Price() float64 {
              return p.price
            }

            func (p *Product) SetPrice(newPrice float64) {
              p.price = newPrice
            }

49. defining constructor 

            func NewProduct(name, category string, price float64) *Product {
             return &Product{ name, category, price }
            }

50. composition : reusing structs 

            type Product struct {
              name string
              *Price  //check without pointers
            }

            type Price struct {
              amount int
              tax    int
            }

            func main() {
              var price *Price = &Price{amount: 100, tax: 5}
              var name string = "mobile"
              var product = Product{name, price}
              fmt.Println(product)
            }

60. map test

              package main

              import (
                "golang.org/x/tour/wc"
                "fmt"
                "strings"
              )

              func WordCount(s string) map[string]int {
                // creating map
                var m = make(map[string]int)
                // split words from string and store inside slic
                words := strings.Split(s, " ")
                // iterate words 
                for _, word := range words {
                  // check if word already exists, if yes increment, if no set 1(first occurance)
                  if _, ok := m[word]; ok {
                    m[word] = m[word] + 1
                  } else {
                    m[word] = 1
                  }
                }
                fmt.Println(m)
                return m
              }

              func main() {
                wc.Test(WordCount)
              }

61. passing function as value

              package main

              import (
                "fmt"
              )

              func computer(square func(int) int, value int) int {
                return square(value)
              }

              func main() {
                var f func(int) int = func(x int) int {
                  return x * x
                }
                fmt.Printf("%T \n", f)
                fmt.Println(computer(f, 5))
              }
  
62. closure 

                package main

                import (
                  "fmt"
                )

                func incrementor() func() int {
                  var count int = 0
                  return func() int {
                    count++
                    return count
                  }
                }

                func main() {
                  var inc func() int = incrementor()
                  fmt.Println(inc())
                  fmt.Println(inc())
                  fmt.Println(inc())
                  fmt.Println(inc())
                  fmt.Println(inc())
                }
63. fibonocci using closure

                  package main

                  import "fmt"

                  // fibonacci is a function that returns
                  // a function that returns an int.
                  func fibonacci() func() int {
                    a,b := 0,1
                    var flagOnce bool = true
                    state := 0
                    return func() int {
                      if(flagOnce) {
                        flagOnce = false
                        return 0
                      }
                      a = b
                      b = state
                      state = a + b
                      return state
                    }
                  }

                  func main() {
                    f := fibonacci()
                    for i := 0; i < 10; i++ {
                      fmt.Println(f())
                    }
                  }

 
64. methods created on builtin datatypes like struct

                package main

                import (
                  "fmt"
                )

                type mynumber int

                func (p mynumber) square() {
                  fmt.Println(p * p)
                }

                func main() {
                  mynumber.square(7)
                }
            
