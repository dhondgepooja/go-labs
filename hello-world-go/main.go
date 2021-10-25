package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {

	// How to declare a variable Option 1:
	var message1 string
	message1 = "hello from "

	//Option 2:
	message2 := "the other side"
	sayTheMessage(message1 + message2)

	// How to call a method from a package
	whatElizaSays := doctor.Intro()
	sayTheMessage(whatElizaSays)

	// How to read user input
	reader := bufio.NewReader(os.Stdin)

	// How to use loop
	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1) // handle new line for windows
		userInput = strings.Replace(userInput, "\n", "", -1)   // handle new line for mac/linux

		// userInput == "quit" will not work because we have \n after that
		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}
	}

}

func sayTheMessage(whatToSay string) {
	fmt.Println(whatToSay)
}
