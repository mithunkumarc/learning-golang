  #### function syntax : 

        func receiver identifier(parameters) returnType {
          //body
        }
        
        // receiver : optional

#### curly braces must start at the end of function declaration

        // error
        func main() 
        {
          fmt.Println("Hello, playground")
        }
        
        // ok
        func main() {
	        fmt.Println("Hello, playground")
        }

