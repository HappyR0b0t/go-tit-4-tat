package game

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
