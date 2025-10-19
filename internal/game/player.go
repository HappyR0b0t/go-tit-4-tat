package game

type Player struct {
	playerName  string
	playerMove  int
	playerScore int
	Strategy    Strategy
	totalScore  int
}

func (p Player) MakeMove(turnNumber int) int {
	return p.Strategy.Move(turnNumber)
}

func NewPlayer(playerName string, strategy Strategy) *Player {
	return &Player{
		playerName:  playerName,
		playerMove:  0,
		playerScore: 0,
		Strategy:    strategy,
		totalScore:  0,
	}
}
