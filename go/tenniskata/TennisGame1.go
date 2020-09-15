package tenniskata

type tennisGame1 struct {
	scorePlayer1 int
	scorePlayer2 int
	player1Name  string
	player2Name  string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	game := &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name}

	return game
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.scorePlayer1 += 1
	} else {
		game.scorePlayer2 += 1
	}
}

func (game *tennisGame1) GetScore() string {
	if game.tied() {
		return tiedScore(game)
	}

	if game.winOrAdvantage() {

		differenceInPoints := game.scorePlayer1 - game.scorePlayer2
		return winOrAdvantage(differenceInPoints) + leadingPlayerAsString(differenceInPoints)
	}

	return stringValueFor(game.scorePlayer1) + "-" + stringValueFor(game.scorePlayer2)
}

func leadingPlayerAsString(differenceInPoints int) string {
	if differenceInPoints < 0 {
		return " player2"
	}
	return " player1"
}

func winOrAdvantage(points int) string {
	if points >= 2 || points <= -2 {
		return "Win for"
	}
	return "Advantage"
}

func (game *tennisGame1) winOrAdvantage() bool {
	return game.scorePlayer1 >= 4 || game.scorePlayer2 >= 4
}

func tiedScore(game *tennisGame1) string {
	switch game.scorePlayer1 {
	case 0:
		return "Love-All"
	case 1:
		return "Fifteen-All"
	case 2:
		return "Thirty-All"
	default:
		return "Deuce"
	}
}

func (game *tennisGame1) tied() bool {
	return game.scorePlayer1 == game.scorePlayer2
}

func (game *tennisGame1) winForPlayerOne(differenceInPoints int) bool {
	return differenceInPoints >= 2
}

func (game *tennisGame1) advantageForPlayerTwo(differenceInPoints int) bool {
	return differenceInPoints == -1
}

func (game *tennisGame1) advantageForPlayerOne(differenceInPoints int) bool {
	return differenceInPoints == 1
}

func stringValueFor(score int) string {
	switch score {
	case 0:
		return "Love"
	case 1:
		return "Fifteen"
	case 2:
		return "Thirty"
	default:
		return "Forty"
	}
}
