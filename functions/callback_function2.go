package main

import (
	"fmt"
)

func main() {
	//callback function : sending function as argument
	login_name := "vichitra" // success
	// login_name := "" //failure
	message := login(validation, login_name)
	fmt.Println(message)
}

//this function is being sent as argument to login
func validation(name string) bool {
	if(len(name) > 0){
		return true
	}
	return false
}

//first arugment is validating function
//validation function signature must be same, arugments and return type
func login(f func(name string) bool , name string) string {
	t := f(name)
	if(t){
		return "success"
	}
	return "failure"
}
