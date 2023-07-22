#### Note

  Go doesn't have char data type, use string with single char for gender type M/F, boolean Y/N similar to javascript

#### source : https://www.callicoder.com/golang-basic-types-operators-type-conversion/

#### Golang has two additional integer types called byte and rune that are aliases for uint8 and int32 data types respectively -

      Type	Alias For
      byte	uint8
      rune	int32

1. In Go, the byte and rune data types are used to distinguish characters from integer values.
2. Golang doesn’t have a char data type. It uses byte and rune to represent character values.
3. The byte data type represents ASCII characters and the rune data type represents a more broader set of Unicode characters that are encoded in UTF-8 format.
4. Characters are expressed in Golang by enclosing them in single quotes like this: 'A'.
5. The default type for character values is rune. That means, if you don’t declare a type explicitly when declaring a variable with a character value, then Go will infer the type as rune -

        
        var firstLetter = 'A' // Type inferred as `rune` (Default type for character values)
        You can create a byte variable by explicitly specifying the type -
        
        var lastLetter byte = 'Z'
        Both byte and rune data types are essentially integers.
        For example, a byte variable with value 'a' is converted to the integer 97.
        
        a rune variable with a unicode value '♥' is converted to the corresponding unicode codepoint U+2665,
        where U+ means unicode and the numbers are hexadecimal, which is essentially an integer.

#### byte vs rune 

        package main
        import "fmt"
        
        func main() {
        	var myByte byte = 'a'
        	var myRune rune = '♥'
        
        	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)
        }
        # Output
        a = 97 and ♥ = U+2665
        In the above example, I’ve printed the variable myByte in character and decimal format, 
        and the variable myRune in character and Unicode format.
