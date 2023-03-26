package tournamentwinner

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestTournamentWinnerProbSuite struct {
	suite.Suite
	TournamentWinnerProb
}

func TestMain(t *testing.T) {
	suite.Run(t, new(TestTournamentWinnerProbSuite))
}

func (suite *TestTournamentWinnerProbSuite) SetupTest() {
	suite.TournamentWinnerProb = NewTournamentWinnerProb()
}

func (s *TestTournamentWinnerProbSuite) Test_ShouldCallAndReturnTournamentWinner() {
	teams := [][]string{{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"}}
	actual := s.TournamentWinnerProb.TournamentWinner(teams, []int{0, 0, 1})
	s.Suite.NotNil(actual)
	s.Suite.Equal("Python", actual)

}
