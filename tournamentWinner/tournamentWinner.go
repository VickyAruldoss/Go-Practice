package tournamentwinner

type TournamentWinnerProb interface {
	TournamentWinner(competitions [][]string, results []int) string
}

type tournamentWinnerProb struct {
}

func NewTournamentWinnerProb() *tournamentWinnerProb {
	return &tournamentWinnerProb{}
}

func (t *tournamentWinnerProb) TournamentWinner(competitions [][]string, results []int) string {
	points := map[string]int{}
	winner := ""
	for index, currentValue := range competitions {
		w := currentValue[1-results[index]]
		points[w] += 3
		if points[w] > points[winner] {
			winner = w
		}
	}
	return winner
}
