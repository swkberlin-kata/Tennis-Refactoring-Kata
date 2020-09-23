package tenniskata

type tennisGame1 struct {
	player1Points int
	player2Points int
	player1Name   string
	player2Name   string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	return &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name,
	}
}

var pointsToScore = map[int]string{
	0: "Love",
	1: "Fifteen",
	2: "Thirty",
	3: "Forty",
}

func (game *tennisGame1) GetScore() string {
	if game.tied() {
		return tiedScore(game)
	}

	if game.winOrAdvantage() {
		diff := differenceInPoints(game)

		return winOrAdvantageFor(diff) + playerLeadingBasedOn(diff)
	}

	return pointsToScore[game.player1Points] + "-" + pointsToScore[game.player2Points]
}

func (game *tennisGame1) tied() bool {
	return game.player1Points == game.player2Points
}

func (game *tennisGame1) winOrAdvantage() bool {
	return game.player1Points >= 4 || game.player2Points >= 4
}

func differenceInPoints(game *tennisGame1) int {
	return game.player1Points - game.player2Points
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.player1Points += 1
	} else {
		game.player2Points += 1
	}
}

func playerLeadingBasedOn(differenceInPoints int) string {
	if differenceInPoints >= 0 {
		return " player1"
	} else {
		return " player2"
	}
}

func winOrAdvantageFor(points int) string {
	if points >= 2 || points <= -2 {
		return "Win for"
	}
	return "Advantage"
}

func tiedScore(game *tennisGame1) string {
	switch game.player1Points {
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
