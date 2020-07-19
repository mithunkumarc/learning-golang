no public private in golang, just exporing or not exporing  

export :

        if you declare var/func with Uppercase you are exporting that var/fun

        eg: "fmt" package Println is starts with uppercase so we can access  

        user defined : 
        
            func Factorial() {
             //...
            }

not export/private :

        if you declare var/func with lowercase you are making it private
        
          func factorial() {
            //..
          }
