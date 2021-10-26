package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

var reader = bufio.NewReader(os.Stdin)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber    int
	PlayerScore    int
	CompluterScore int
}

func (g *Game) Rounds() {
	// use select to process input in channels
	// print to screen
	// keep track of round number
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
			// Option 2 using channel
			//g.DisplayChan <- ""
		}
	}
}

// clearScreen clears the screen
func (g *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *Game) PrintIntro() {
	// Option 1
	fmt.Println("Rock, Paper and Scissors")
	fmt.Println("------------------------")
	fmt.Println("Game is played in 3 rounds. Best of 3 wins")

	// Option 2 with channel
	/*g.DisplayChan <- fmt.Sprintf(`
	Rock, Paper and Scissors
	------------------------
	Game is played in 3 rounds. Best of 3 wins
	`)
		<-g.DisplayChan*/ // This is done so that we wait for the channel to finish handling what we just sent above. The channel should send something once it is done hence we add line while printing msg
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1

	fmt.Println()
	fmt.Println("Round", g.Round.RoundNumber)
	fmt.Println("---------------")
	fmt.Print("Please enter rock, paper, or scissors ->")

	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	// If else
	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	fmt.Println()
	g.DisplayChan <- fmt.Sprintf("Player Chose %s", playerChoice)

	// Switch
	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose rock")
		break
	case PAPER:
		fmt.Println("Computer chose paper")
		break
	case SCISSORS:
		fmt.Println("Computer chose scissors")
		break
	default:
	}

	if playerValue == computerValue {
		g.DisplayChan <- "It's a draw"
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		default:
			g.DisplayChan <- "Player chose invalid."
			return false
		}
	}
	return true
}

func (g *Game) playerWins() {
	g.DisplayChan <- "Player Wins"
	g.Round.PlayerScore++
}

func (g *Game) computerWins() {
	g.DisplayChan <- "Computer Wins"
	g.Round.CompluterScore++
}

func (g *Game) PrintSummary() {
	fmt.Println("Final Score")
	fmt.Println("------------")
	fmt.Printf("Player Score %d/3, Computer Score %d/3", g.Round.PlayerScore, g.Round.CompluterScore)
	fmt.Println()
	if g.Round.PlayerScore > g.Round.CompluterScore {
		fmt.Println("Player wins.")
	} else {
		fmt.Println("Computer wins.")
	}
}
