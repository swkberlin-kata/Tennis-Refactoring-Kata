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

type game1 struct {
	player1, player2 game1Player
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	return &game1{
		player1: game1Player{ name: player1Name },
		player2: game1Player{ name: player2Name },
	}
}

func (g *game1) GetScore() string {
	if g.player1.points == g.player2.points {
		return tiedScore(g)
	}

	if g.winOrAdvantage() {
		diff := differenceInPoints(g)

		return winOrAdvantageFor(diff) + " " + g.leadingPlayer()
	}

	return g.player1.Score() + "-" + g.player2.Score()
}

func (g *game1) winOrAdvantage() bool {
	return g.player1.points >= 4 || g.player2.points >= 4
}

func differenceInPoints(g *game1) int {
	return g.player1.points - g.player2.points
}

func (g *game1) WonPoint(playerName string) {
	if playerName == "player1" {
		g.player1.points += 1
	} else {
		g.player2.points += 1
	}
}

// will not return proper if equal.
func (g *game1) leadingPlayer() string {
	if g.player1.points >= g.player2.points {
		return g.player1.name
	}

	return g.player2.name
}

func winOrAdvantageFor(points int) string {
	if points >= 2 || points <= -2 {
		return "Win for"
	}
	return "Advantage"
}

func tiedScore(g *game1) string {
	switch g.player1.points {
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
