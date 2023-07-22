FPrintf : prints and It returns the number of bytes written and any write error encountered.


    // Declaring some const variables
    const name, dept = "GeeksforGeeks", "CS"
 
    // Calling Fprintf() function which returns
    // "n" as the number of bytes written and
    // "err" as any error ancountered
    n, err := fmt.Fprintf(os.Stdout, "%s is a %s portal.\n",
                                                name, dept)
 
    // Printing the number of bytes written
    fmt.Print(n, " bytes written.\n")
 
    // Printing if any error encountered
    fmt.Print(err)

    Output: 
 
    
    GeeksforGeeks is a CS portal.
    30 bytes written.
    <nil>

SPrintf : Save Formatted string and returns it 

     const name, dept = "GeeksforGeeks", "CS"
  
    // Calling Sprintf() function
    s := fmt.Sprintf("%s is a %s Portal.\n", name, dept)
  
    // Calling WriteString() function to write the
    // contents of the string "s" to "os.Stdout"
    io.WriteString(os.Stdout, s)
  
    
    Output:    
    GeeksforGeeks is a CS Portal.
