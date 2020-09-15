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
	score := ""
	if game.scorePlayer1 == game.scorePlayer2 {
		switch game.scorePlayer1 {
		case 0:
			score = "Love-All"
		case 1:
			score = "Fifteen-All"
		case 2:
			score = "Thirty-All"
		default:
			score = "Deuce"
		}
		return score
	}

	if game.scorePlayer1 >= 4 || game.scorePlayer2 >= 4 {
		differenceInPoints := game.scorePlayer1 - game.scorePlayer2
		if game.advantageForPlayerOne(differenceInPoints) {
			score = "Advantage player1"
			return score
		}
		if game.advantageForPlayerTwo(differenceInPoints) {
			score = "Advantage player2"
			return score
		}

		if game.winForPlayerOne(differenceInPoints) {
			score = "Win for player1"
			return score
		}

		return "Win for player2"
	}

	return stringValueFor(game.scorePlayer1) + "-" + stringValueFor(game.scorePlayer2)
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
	case 3:
		return "Forty"
	}
	return ""
}
