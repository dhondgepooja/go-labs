package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

// User defined data structure
type User struct {
	Name            string
	Age             int
	FavouriteNumber float64
	OwnsADog        bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User
	user.Name = readString("What is your name?")
	user.Age = readInt("What is your age?")
	user.FavouriteNumber = readFloat("What is your favourite number?")
	user.OwnsADog = readBool("Do you have a dog? (y/n)")

	// OUTPUT: "Your name is pooja . and You are 32 years old."
	//fmt.Println("Your name is", name, ". You are", age, "years old.")             // NOTE: this will print with weird space "Your name is pooja . and You are 32 years old."

	// OUTPUT: "Your name is pooja. You are 32 years old."
	//fmt.Println("Your name is "+name+". You are", age, "years old.")              // SLOWER and requires more memory also can't do concat for age since it is an int
	//fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.", name, age)) // Sprintf returns formatted string

	fmt.Printf("Your name is %s. You are %d years old. Your favourite number is %.4f. Owns a dog: %t \n", user.Name, user.Age, user.FavouriteNumber, user.OwnsADog)

}

func prompt() {
	fmt.Print("->")
}

func readString(question string) string {
	for {
		fmt.Println(question)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		userInput = strings.Replace(userInput, "\r\n", "", -1) // NOTE: keep this one in mind. You may forget handling this for windows

		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(question string) int {
	for {
		fmt.Println(question)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat(question string) float64 {
	for {
		fmt.Println(question)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readBool(question string) bool {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(question)
		prompt()

		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(char))
		if char == 'y' || char == 'Y' {
			return true
		} else if char == 'n' || char == 'N' {
			return false
		} else {
			fmt.Println("Enter y/n")
		}

	}
}
