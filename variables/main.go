package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready"

func main() {
	// option 1 var
	//var firstNumber int
	// firstNumber = 2

	// option 2 var
	//var secondNumber = 5

	// option 3 var
	//subtraction := 7

	rand.Seed(time.Now().UnixNano()) // this will ensure all 2 random numbers a are distinct

	var firstNumber = rand.Intn(8) + 2 // this will ensure firstNumber is between 2 and 10
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction // defaults to 0

	reader := bufio.NewReader(os.Stdin)

	playGame(reader, firstNumber, secondNumber, subtraction, answer)
}

func playGame(reader *bufio.Reader, firstNumber, secondNumber, subtraction, answer int) {
	//display welcome/instruction
	fmt.Println("Guess the Number Game")
	fmt.Println("----------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	//take them through the games
	//fmt.Println("Multiply your number by", firstNumber, "and press enter when ready") //NOTE: automatically adds space
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiple the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("The answer is", answer)
}
