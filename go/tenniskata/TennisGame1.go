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
		return winOrAdvantageFor(differenceInPoints(game)) + playerLeadingBasedOn(differenceInPoints(game))
	}

	return pointsToScore[game.scorePlayer1] + "-" + pointsToScore[game.scorePlayer2]
}

func (game *tennisGame1) tied() bool {
	return game.scorePlayer1 == game.scorePlayer2
}

func (game *tennisGame1) winOrAdvantage() bool {
	return game.scorePlayer1 >= 4 || game.scorePlayer2 >= 4
}

func differenceInPoints(game *tennisGame1) int {
	return game.scorePlayer1 - game.scorePlayer2
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.scorePlayer1 += 1
	} else {
		game.scorePlayer2 += 1
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
