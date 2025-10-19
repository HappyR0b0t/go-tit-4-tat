package main

import (
	"fmt"

	"example.com/go-tit-4-tat/internal/game"
)

// --- Match ---

// --- PLayer ---

// --- Strategies ---

// --- Tournament ---

// --------------------------------------------------------------------------

// --------------------------------------------------------------------------
func main() {
	fmt.Println("Hello, this is tit-4-tat")

	// match := Match{}

	// Players setup
	// playerOne := NewPlayer("One", AlwaysOneStrategy{})
	// playerTwo := NewPlayer("Two", RandomStrategy{})
	playerNames := []string{
		"One",
		"Two",
		"Three",
		"Four"}
	strategies := []game.Strategy{
		game.AlwaysOneStrategy{},
		game.AlwaysZeroStrategy{},
		game.RandomStrategy{},
		game.OnesAndZeroesStrategy{}}

	players := []game.Player{}
	for i := range playerNames {
		players = append(players, *game.NewPlayer(playerNames[i], strategies[i]))
	}

	// // Match setup
	numberOfTurns := 20
	// match := NewMatch(*playerOne, *playerTwo, numberOfTurns)
	// match.InitTurns(numberOfTurns)

	// // Match play
	// match.PlayMatch(playerOne, playerTwo)
	// fmt.Println(match)
	// fmt.Println("The rusult of match is:")
	// fmt.Println("Player one score is:", match.playerOneScore, "Player two score is:", match.playerTwoScore)

	match := game.Match{}

	//Tournament setup

	// players := []Player{*playerOne, *playerTwo}
	matches := match.MakeMatchesV2(players, numberOfTurns)
	fmt.Println(matches)
	tournament := game.NewTournament(len(players), 1, numberOfTurns, matches)
	tournament.StartTournament()
	fmt.Println(tournament.Highscore)
}
