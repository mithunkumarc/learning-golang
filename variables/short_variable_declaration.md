#### short variable declarations


      A varialbe can be declared and assigned togother using short variable declaration opeartor
      the operator uses this symbol  :=
      
      example:
      
      count := 100
      
> Note: short variable declaration works only inside function, not outside function(package scope)

      example:
      
      color := "green"    // syntax error: unexpected :=, expecting =
      func ScanInput() {
        //..
      }
