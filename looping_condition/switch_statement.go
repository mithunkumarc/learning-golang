package main

import (
	"fmt"
)

/* break statement not needed after each case,  its automatically added*/

func getData(data string) string {
	switch data {
	case "monday":
		return "monday"
	case "tuesday":
		return "tuesday"
	default:
		return "other day"
	}
}

func main() {

	var flag string = "monday"
	switch flag {
	case "monday":
		fmt.Println("monday")
	case "tuesday":
		fmt.Println("tuesday")
	default:
		fmt.Println("other day")
	}

	result := getData("sunday") // return function
	fmt.Println(result) // other day

	/* fallthrough used to execute next immediate case*/

	var product string = "fmcg"
	switch product {
	case "electronics":
		fallthrough
	case "mobile":
		fmt.Println("mobile")
		fallthrough
	case "tv":
		fmt.Println("tv")
	case "fmcg":
		fallthrough
	case "soap":
		fmt.Println("soap")
		fallthrough
	case "toothpaste":
		fmt.Println("toothpaste")
	default:
		fmt.Println("none")
	}

}
