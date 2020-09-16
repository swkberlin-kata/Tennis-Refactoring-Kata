package tenniskata

type tennisGame1 struct {
	m_score1    int
	m_score2    int
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
	if playerName == game.player1Name {
		game.m_score1 += 1
	} else {
		game.m_score2 += 1
	}
}

var game1ScoreToName = map[int]string{
	0: "Love",
	1: "Fifteen",
	2: "Thirty",
}

func (game *tennisGame1) GetScore() string {
	if game.m_score1 == game.m_score2 {
		// There are special names for each of the scores where they are tied.

		if name, ok := game1ScoreToName[game.m_score1]; ok {
			return name + "-All"
		}

		return "Deuce"
	}

	if game.m_score1 >= 4 || game.m_score2 >= 4 {
		minusResult := game.m_score1 - game.m_score2

		if minusResult == 1 {
			return "Advantage player1"
		}

		if minusResult == -1 {
			return "Advantage player2"
		}

		if minusResult >= 2 {
			return "Win for player1"
		}

		return "Win for player2"
	}

	var score string

	if name, ok := game1ScoreToName[game.m_score1]; ok {
		score += name
	} else {
		score += "Forty"
	}

	score += "-"

	if name, ok := game1ScoreToName[game.m_score2]; ok {
		score += name
	} else {
		score += "Forty"
	}

	return score
}
