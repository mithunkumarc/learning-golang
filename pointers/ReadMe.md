#### there is a difference in dereferencing(reading value) of primitive values and object(struct) reference variable

use * to deference primitive value

        t := 10
        s := &t
        fmt.Println(*s) // if u skip *, u will get address

no need to use * for object(struct) reference variable to read state,  because it is already pointing to memory locaiton


          p := Person{name: "raj", age: 12}
          q := &p         // q is a pointer which doesn't require * to dereference(read)
          q.name = "arun" // changing state
          fmt.Println(q)	// &{arun 12}
          fmt.Println(*q) // {arun 12} for reading whole object is fine     
          fmt.Println(q.name) // arun , * is not used to access state
