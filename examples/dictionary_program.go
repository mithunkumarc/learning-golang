package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func SeparatorLine() {
	fmt.Println(strings.Repeat("*", 10))
}

func PrintMessage()  {
	fmt.Println("Enter below options");
	fmt.Println("Enter 1 to add");
	fmt.Println("Enter 2 to delete");
	fmt.Println("Enter 3 to update");
	fmt.Println("Enter 4 to search");
	fmt.Println("Enter 5 to print all words")
	fmt.Println("Enter 6 to exit");
}

func main() {
	dictionary := make(map[string]string)
	for true {
		PrintMessage()
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')		
		result, _ := strconv.ParseUint(input[:len(input)-1], 10, 64)		
		switch result {
			case 1:
				fmt.Println("You have entered option 1")
				fmt.Println("Enter Word to add")				
				addWReader := bufio.NewReader(os.Stdin)
				// read word
				addWInput, _ := addWReader.ReadString('\n')
				fmt.Println("Enter Meaning to add")
				addMReader := bufio.NewReader(os.Stdin)
				// read meaning
				addMInput, _ := addMReader.ReadString('\n')
				addWInput = addWInput[:len(addWInput)-1]
				addMInput = addMInput[:len(addMInput)-1]
				fmt.Println(addWInput, addMInput)
				dictionary[addWInput] = addMInput
				msg := fmt.Sprintf("%s added world successfully", addWInput)
				fmt.Println(msg)
				continue
			case 2:
				fmt.Println("You have entered option 2")
				fmt.Println("Enter word to delete")
				delWReader := bufio.NewReader(os.Stdin)
				// read word
				delWinput, _ := delWReader.ReadString('\n')
				delWinput = delWinput[:len(delWinput)-1]
				_, found := dictionary[delWinput]
				if found {
					delete(dictionary, delWinput)
					fmt.Println(delWinput, " deleted successfully")
				} else {
					fmt.Println("word not found")
				}
				continue
			case 3:
				fmt.Println("You have entered option 3")
				fmt.Println("Enter word to update")
				updateWReader := bufio.NewReader(os.Stdin)
				// read word
				updateWinput, _ := updateWReader.ReadString('\n')
				updateWinput = updateWinput[:len(updateWinput)-1]
				_, found := dictionary[updateWinput]
				if found {
					fmt.Println("Enter new meaning for ",updateWinput)
					updateMinput, _ := updateWReader.ReadString('\n')
					updateMinput = updateMinput[:len(updateMinput)-1]
					dictionary[updateWinput] = updateMinput
					fmt.Println(updateWinput, " word updated successfully")
				} else {
					fmt.Println(updateWinput, " word not found, not able to update")
				}
				continue
			case 4:
				fmt.Println("You have entered option 4")
				fmt.Println("Enter word to search")
				searchWReader := bufio.NewReader(os.Stdin)
				// read word
				searchWinput, _ := searchWReader.ReadString('\n')
				searchWinput = searchWinput[:len(searchWinput)-1]
				_, found := dictionary[searchWinput]
				if found {
					fmt.Println(searchWinput, " exists!!!")
					fmt.Println(searchWinput,":",dictionary[searchWinput])
				} else {
					fmt.Println(searchWinput, " not found")
				}
				continue
			case 5:
				fmt.Println("You have entered option 5")
				fmt.Println("Reading all words")
				SeparatorLine()
				for w, m := range dictionary {
					fmt.Println(w,":",m)
				}
				SeparatorLine()
				continue
			case 6:
				fmt.Println("Exiting....")
				fmt.Println("Thank You for using my dictionary program")
				return
			default:
				fmt.Println("invalid option")				
				continue
			}		
	}	
}
