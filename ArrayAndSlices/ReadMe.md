
#### Array : specify number of elements while declaration  
#### slice : number of elements not specified while declaration  

#### Array vs Slice Declaration
      
      var intArray = [10] int {0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
      var slice1 [] int = intArray[2:5]  // index 2 to 4, size = 3


Array is fixed in size, cannot be appended(fixed).  
We can append new element to slice(but slice create new slice of array with new and old elements).  
 
#### slice can also be created using existing array using range

    array[start: end]

> slice manipulation

       NewSclice = append(oldSlice, newElement) //immutable
