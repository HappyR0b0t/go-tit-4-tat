package game

type Match struct {
	playerOne      Player
	playerTwo      Player
	winner         Player
	loser          Player
	turnsNumber    int
	turns          [][2]int
	playerOneScore int
	playerTwoScore int
}

func NewMatch(playerOne, playerTwo Player, turnsNumber int) *Match {
	return &Match{
		playerOne:      playerOne,
		playerTwo:      playerTwo,
		winner:         Player{},
		loser:          Player{},
		playerOneScore: 0,
		playerTwoScore: 0,
		turnsNumber:    turnsNumber,
		turns:          [][2]int{},
	}
}

// Create a series of matches between each and each players
func (m Match) MakeMatchesV2(players []Player, numberOfTurns int) []Match {
	matches := []Match{}
	for i := 0; i <= len(players)-1; i++ {
		for j := i; j <= len(players)-1; j++ {
			matches = append(matches, *NewMatch(players[i], players[j], numberOfTurns))
		}
	}
	return matches
}

// Score distribution at the end of the turn
func (m *Match) ScoreDistrib() {
	if m.playerOne.playerMove == 1 && m.playerTwo.playerMove == 1 {
		m.playerOne.playerScore = 3
		m.playerTwo.playerScore = 3
	} else if m.playerOne.playerMove == 0 && m.playerTwo.playerMove == 0 {
		m.playerOne.playerScore = 1
		m.playerTwo.playerScore = 1
	} else if m.playerOne.playerMove == 1 && m.playerTwo.playerMove == 0 {
		m.playerOne.playerScore = 0
		m.playerTwo.playerScore = 5
	} else if m.playerOne.playerMove == 0 && m.playerTwo.playerMove == 1 {
		m.playerOne.playerScore = 5
		m.playerTwo.playerScore = 0
	}
}

func (m *Match) PlayMatch() {
	m.turns = make([][2]int, m.turnsNumber)
	for i := range m.turnsNumber {
		m.playerOne.playerMove = m.playerOne.MakeMove(i)
		m.playerTwo.playerMove = m.playerTwo.MakeMove(i)
		m.turns[i] = [2]int{m.playerOne.playerMove, m.playerTwo.playerMove}

		m.ScoreDistrib()
		m.playerOneScore += m.playerOne.playerScore
		m.playerTwoScore += m.playerTwo.playerScore

		// fmt.Println("This is move number", i+1)
		// fmt.Println("Player one move is:", m.playerOne.playerMove, "Player two move is:", m.playerTwo.playerMove)
		// fmt.Println("Player one score is:", m.playerOneScore, "Player two score is:", m.playerTwoScore)
	}

	m.playerOne.totalScore += m.playerOneScore
	m.playerTwo.totalScore += m.playerTwoScore

	if m.playerOneScore > m.playerTwoScore {
		m.winner = m.playerOne
		m.loser = m.playerTwo
		// fmt.Println("Player One Wins!")
	} else if m.playerOneScore < m.playerTwoScore {
		m.winner = m.playerTwo
		m.loser = m.playerOne
		// fmt.Println("Player Two Wins!")
	} else {
		// fmt.Println("It's a Tie!")
	}
}
