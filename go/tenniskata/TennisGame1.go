package tenniskata

type tennisGame1 struct {
	mScore1     int
	mScore2     int
	player1Name string
	player2Name string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	return &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name,
	}
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.mScore1 += 1
		return
	}
	game.mScore2 += 1
}

func (game *tennisGame1) GetScore() string {
	if game.mScore1 == game.mScore2 {
		switch game.mScore1 {
		case 0:
			return "Love-All"
		case 1:
			return "Fifteen-All"
		case 2:
			return "Thirty-All"
		}
		return "Deuce"
	} else if game.mScore1 >= 4 || game.mScore2 >= 4 {
		minusResult := game.mScore1 - game.mScore2
		if minusResult == 1 {
			return "Advantage player1"
		} else if minusResult == -1 {
			return "Advantage player2"
		} else if minusResult >= 2 {
			return "Win for player1"
		}
		return "Win for player2"
	}
	score := ""
	tempScore := 0
	for i := 1; i < 3; i++ {
		if i == 1 {
			tempScore = game.mScore1
		} else {
			score += "-"
			tempScore = game.mScore2
		}
		switch tempScore {
		case 0:
			score += "Love"
		case 1:
			score += "Fifteen"
		case 2:
			score += "Thirty"
		case 3:
			score += "Forty"
		}
	}
	return score
}
