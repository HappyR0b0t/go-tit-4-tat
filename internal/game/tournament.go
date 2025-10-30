package game

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type Tournament struct {
	numberOfPlayers       int
	numberOfGamesPerMatch int
	numberOfTurnsPerMatch int
	matches               []Match
	Highscore             map[string]int
}

func NewTournament(numberOfPlayers, numberOfGamesPerMatch, numberOfTurnsPerMatch int, matches []Match) *Tournament {
	return &Tournament{
		numberOfPlayers:       numberOfPlayers,
		numberOfGamesPerMatch: numberOfGamesPerMatch,
		numberOfTurnsPerMatch: numberOfTurnsPerMatch,
		matches:               matches,
		Highscore:             map[string]int{},
	}
}

func (t *Tournament) StartTournament() {
	for _, match := range t.matches {
		match.PlayMatch()
		t.Highscore[match.playerOne.playerName] += match.playerOneScore
		t.Highscore[match.playerTwo.playerName] += match.playerTwoScore
	}
}

func (t *Tournament) HighScore() [][]string {
	res := [][]string{}

	score := []int{}
	names := []string{}

	for k, v := range t.Highscore {
		score = append(score, v)
		names = append(names, k)
	}

	sort.Slice(score, func(i, j int) bool {
		if score[i] > score[j] {
			names[i], names[j] = names[j], names[i]
			return true
		}
		return false
	})

	for i := range names {
		res = append(res, []string{"name: " + names[i] + " " + "score: " + strconv.Itoa(score[i])})
	}
	fmt.Println("This is highscore after all matches:")

	if len(res) > 5 {
		return res[:5]
	}
	return res
}

func PlayTournament(numberOfTurns int) {
	fmt.Println("STARTING TOURNAMENT:")
	start := time.Now()

	genericStrategiesTurns := [][]int{}
	BuildCombinations(&genericStrategiesTurns, []int{}, []int{0, 1}, numberOfTurns)
	genericPlayersNames := GenerateNames(len(genericStrategiesTurns))

	players := []Player{}

	genericStrategies := []Strategy{}

	for _, t := range genericStrategiesTurns {
		genericStrategies = append(genericStrategies, GenerateGenericStrategies(GenericStrategy{}, t))
	}

	for i := range genericPlayersNames {
		players = append(players, *NewPlayer(genericPlayersNames[i], genericStrategies[i]))
	}

	fmt.Println("NUMBER OF PLAYERS: ", len(players))

	match := Match{}

	matches := match.MakeMatchesV2(players, numberOfTurns)
	fmt.Println("NUMBER OF MATCHES: ", len(matches))
	tournament := NewTournament(len(players), 1, numberOfTurns, matches)
	tournament.StartTournament()
	fmt.Println(tournament.HighScore())

	duration := time.Since(start)
	seconds := duration.Seconds()
	fmt.Println("Time elapsed: ", seconds, "seconds")
}
