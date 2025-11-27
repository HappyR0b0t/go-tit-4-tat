package main

import (
	"fmt"
	"sync"
	"time"

	"example.com/go-tit-4-tat/internal/game"
)

var numberOfTurns = 8

func main() {
	fmt.Println("Hello, this is tit-4-tat")
	start := time.Now()

	// match := Match{}

	// Players setup
	// playerOne := NewPlayer("One", AlwaysOneStrategy{})
	// playerTwo := NewPlayer("Two", RandomStrategy{})
	// playerNames := []string{
	// 	"Always cooperate",
	// 	"Always defect",
	// 	"Randomly defect and cooperate",
	// 	"One turn cooperate, one turn defect",
	// }

	// strategies := []game.Strategy{
	// 	game.AlwaysOneStrategy{},
	// 	game.AlwaysZeroStrategy{},
	// 	game.RandomStrategy{},
	// 	game.OnesAndZeroesStrategy{},
	// }

	// genericStrategiesTurns := [][]int{}
	// game.BuildCombinations(&genericStrategiesTurns, []int{}, []int{0, 1}, numberOfTurns)
	// genericPlayersNames := game.GenerateNames(len(genericStrategiesTurns))

	// players := []game.Player{}
	// genericStrategies := []game.Strategy{}

	// for _, t := range genericStrategiesTurns {
	// 	genericStrategies = append(genericStrategies, game.GenerateGenericStrategies(game.GenericStrategy{}, t))
	// }

	// for i := range genericPlayersNames {
	// 	players = append(players, *game.NewPlayer(genericPlayersNames[i], genericStrategies[i]))
	// }

	// fmt.Println("NUMBER OF PLAYERS: ", len(players))

	// // // Match setup
	// match := game.Match{}

	// //Tournament setup
	// matches := match.MakeMatchesV2(players, numberOfTurns)
	// fmt.Println("NUMBER OF MATCHES: ", len(matches))
	// tournament := game.NewTournament(len(players), 1, numberOfTurns, matches)
	// tournament.StartTournament()
	// fmt.Println(tournament.HighScore())

	// Basic concurrent setup
	var wg sync.WaitGroup
	resultsChan := make(chan [][]string)

	// wg.Add(numberOfTurns)

	for i := range numberOfTurns {
		wg.Add(1)
		go func(turn int) {
			defer wg.Done()
			result := game.PlayTournament(turn)
			resultsChan <- result
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var finalRes [][][]string

	for res := range resultsChan {
		finalRes = append(finalRes, res)
	}
	fmt.Println(finalRes)

	// var allResults [][]string
	// for r := range results {
	// 	allResults = append(allResults, r)
	// }

	// fmt.Println(allResults)

	// game.PlayTournament(11)

	duration := time.Since(start)
	seconds := duration.Seconds()
	fmt.Println("Time elapsed: ", seconds, "seconds")
}
