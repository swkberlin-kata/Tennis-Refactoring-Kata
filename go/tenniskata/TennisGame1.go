package tenniskata

var pointsToScore = map[int]string{
	0: "Love",
	1: "Fifteen",
	2: "Thirty",
	3: "Forty",
}

type game1Player struct{
	points int
	name string
}

func (p *game1Player) Score() string {
	return pointsToScore[p.points]
}

type tennisGame1 struct {
	player1, player2 game1Player
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	return &tennisGame1{
		player1: game1Player{ name: player1Name },
		player2: game1Player{ name: player2Name },
	}
}

func (game *tennisGame1) GetScore() string {
	if game.player1.points == game.player2.points {
		return tiedScore(game)
	}

	if game.winOrAdvantage() {
		diff := differenceInPoints(game)

		return winOrAdvantageFor(diff) + " " + game.leadingPlayer()
	}

	return game.player1.Score() + "-" + game.player2.Score()
}

func (game *tennisGame1) winOrAdvantage() bool {
	return game.player1.points >= 4 || game.player2.points >= 4
}

func differenceInPoints(game *tennisGame1) int {
	return game.player1.points - game.player2.points
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.player1.points += 1
	} else {
		game.player2.points += 1
	}
}

// will not return proper if equal.
func (game *tennisGame1) leadingPlayer() string {
	if game.player1.points >= game.player2.points {
		return game.player1.name
	}

	return game.player2.name
}

func winOrAdvantageFor(points int) string {
	if points >= 2 || points <= -2 {
		return "Win for"
	}
	return "Advantage"
}

func tiedScore(game *tennisGame1) string {
	switch game.player1.points {
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
