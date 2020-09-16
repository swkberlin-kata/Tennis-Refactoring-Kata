package tenniskata

type tennisGame1 struct {
	score1    int
	score2    int
	player1Name string
	player2Name string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	return &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name,
	}
}

// WonPoint is when playerName has scored a point.
func (game *tennisGame1) WonPoint(playerName string) {
	switch playerName {
	case game.player1Name:
		game.score1++
	case game.player2Name:
		game.score2++
	}
}

var game1ScoreToName = map[int]string{
	0: "Love",
	1: "Fifteen",
	2: "Thirty",
	3: "Forty",
}

func (game *tennisGame1) GetScore() string {
	if game.score1 == game.score2 {
		// There are special names for each of the scores where they are tied.

		if score, ok := game1ScoreToName[game.score1]; ok && score != "Forty" {
			// if score is not in map  then score = "", and ok == false
			return score + "-All"
		}

		return "Deuce"
	}

	if game.score1 >= 4 || game.score2 >= 4 {
		// One of the players could win here.

		scoreDifference := game.score1 - game.score2

		if scoreDifference == 1 {
			return "Advantage player1"
		}

		if scoreDifference == -1 {
			return "Advantage player2"
		}

		if scoreDifference >= 2 {
			return "Win for player1"
		}

		return "Win for player2"
	}

	// Otherwise, just name of points for p1, dash name of points for p2

	p1score := game1ScoreToName[game.score1]

	p2score := game1ScoreToName[game.score2]
	

	return p1score + "-" + p2score
}
