package sortedsquaredarray

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSortedSquaredArraySuite struct {
	suite.Suite
	SortedSquaredArrayProb
}

func TestMain(t *testing.T) {
	suite.Run(t, new(TestSortedSquaredArraySuite))
}

func (s *TestSortedSquaredArraySuite) SetupTest() {
	s.SortedSquaredArrayProb = NewSortedSquaredArray()
}

func (s *TestSortedSquaredArraySuite) Test_ShouldCallAndReturnNonArrayValues() {
	expected := []int{1, 4, 9, 16, 25}
	actual := s.SortedSquaredArrayProb.SortedSquaredArray([]int{1, 2, 3, 4, 5})
	s.Suite.Equal(expected, actual)
}
