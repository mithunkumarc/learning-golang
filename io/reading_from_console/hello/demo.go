package hello

import "fmt"

func ScanInput() {
	var msg string
	msg = "Welcome to Golang!!"
	fmt.Printf("Hello World %s \n", msg)

	//accept input from user
	var name string
	fmt.Println("Enter your name : ")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("Hello %s \n", name)

	var number int
	fmt.Println("Enter the number :")
	fmt.Scanf("%d\n", &number)
	fmt.Printf("the number is %d \n", number)
}
