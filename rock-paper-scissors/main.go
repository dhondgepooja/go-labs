package main

import (
	game2 "rockpaperscissorapp/game"
)

func main() {
	displayChan := make(chan string)
	roundChan := make(chan int)

	game := game2.Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Round: game2.Round{
			RoundNumber:    0,
			PlayerScore:    0,
			CompluterScore: 0,
		},
	}
	go game.Rounds()
	game.ClearScreen()
	game.PrintIntro()

	for {
		game.RoundChan <- 1
		<-game.RoundChan // wait for the chan

		if game.Round.RoundNumber > 3 {
			break
		}

		if !game.PlayRound() {
			game.RoundChan <- -1
			<-game.RoundChan
		}
	}

	game.PrintSummary()
}
