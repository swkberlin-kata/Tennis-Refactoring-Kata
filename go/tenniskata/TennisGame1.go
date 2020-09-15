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
		if differenceInPoints == 1 {
			score = "Advantage player1"
		} else if differenceInPoints == -1 {
			score = "Advantage player2"
		} else if differenceInPoints >= 2 {
			score = "Win for player1"
		} else {
			score = "Win for player2"
		}
		return score
	}

	return stringValueFor(game.scorePlayer1) + "-" + stringValueFor(game.scorePlayer2)
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
