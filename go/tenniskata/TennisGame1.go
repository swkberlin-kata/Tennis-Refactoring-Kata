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
	switch playerName {
	case g.player1.name:
		g.player1.points += 1
	case g.player2.name:
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
	score := g.player1.Score()
	if score == "" || score == "Forty" {
		return "Deuce"
	}
	return score + "-All"
}
