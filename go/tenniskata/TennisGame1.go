package tenniskata

var scoreNames = [4]string{"Love", "Fifteen", "Thirty", "Forty"}

type tennisGame1 struct {
	mScore1     int
	mScore2     int
	player1Name string
	player2Name string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	game := &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name}

	return game
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.mScore1 += 1
	} else {
		game.mScore2 += 1
	}
}

func (game *tennisGame1) GetScore() string {
	score := ""
	tempScore := 0
	if game.mScore1 == game.mScore2 {
		if game.mScore1 >= 3 {
			score = "Deuce"
		} else {
			score = scoreNames[game.mScore1] + "-All"
		}
	} else if game.mScore1 >= 4 || game.mScore2 >= 4 {
		minusResult := game.mScore1 - game.mScore2
		if minusResult == 1 {
			score = "Advantage player1"
		} else if minusResult == -1 {
			score = "Advantage player2"
		} else if minusResult >= 2 {
			score = "Win for player1"
		} else {
			score = "Win for player2"
		}
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = game.mScore1
			} else {
				score += "-"
				tempScore = game.mScore2
			}
			score += scoreNames[tempScore]
		}
	}
	return score
}
