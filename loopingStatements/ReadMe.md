braces should begin when looping/conditional declaration end, below code will give error, by default golang puts semicolon at the end of the line.

        for i := 0 ; i < 5 ; i++    // error :  unexpected newline, expecting { after for clause
          {
            fmt.Println(i)
          }
          
fix : 

        for i := 0 ; i < 5 ; i++ {
            fmt.Println(i)
          }

todo

        https://yourbasic.org/golang/for-loop/
