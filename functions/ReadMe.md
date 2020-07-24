  #### function syntax : 

        func receiver identifier(parameters) returnType {
          //body
        }
        
        // receiver : optional
	// if a function is returning value, returnType is mandatory
	
#### curly braces must start at the end of function declaration

        // error, reason: end of line golang compiler puts semicolumns default
        func main() 	// at the end golang compiler puts semicolums and it it think funciton body is missing
        {
          fmt.Println("Hello, playground")
        }
        
        // ok
        func main() {
	        fmt.Println("Hello, playground")
        }

