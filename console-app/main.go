package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"strconv"
)

func main() {
	/*reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("->")

		userInput, _ := reader.ReadString('\n') // _ mean i don't care about second retured arg

		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(userInput)
		}
	}*/

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() { // defer keyword means anything following this keyword will happen at the very end of the function
		_ = keyboard.Close() // _ mean we want to ignore what is returned by keyboard.Close
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("____")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	char := ' '

	for char != 'q' && char != 'Q' {
		char, _, err = keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		//t1 := fmt.Sprintf("You chose %q", char) // char is of type rune which is basically the char java type
		//fmt.Println(t1)                         // this prints 0 as '0'
		fmt.Println(string(char))
		i, _ := strconv.Atoi(string(char))

		// OPTION 1
		_, ok := coffees[i]
		if ok {
			t3 := fmt.Sprintf("you chose %s", coffees[i])
			fmt.Println(t3)
		}

		// OPTION 2
		if _, ok := coffees[i]; ok {
			t3 := fmt.Sprintf("you chose %s", coffees[i])
			fmt.Println(t3)
		}

		//t2 := fmt.Sprintf("You chose %d", i) // this prints 0  as 0
		//fmt.Println(t2)

	}

	fmt.Println("Program exiting")
}
