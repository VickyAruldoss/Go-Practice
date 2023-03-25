package twonumbersum

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TwoNumberSumCalculatorSuite struct {
	suite.Suite
	TwoNumberSumCalculator
}

func TestTwoNumberSumCalculatorSuite(t *testing.T) {
	suite.Run(t, new(TwoNumberSumCalculatorSuite))
}

func (suite *TwoNumberSumCalculatorSuite) SetupTest() {
	suite.TwoNumberSumCalculator = NewTwoNumberSumCalculator()
}

func (s *TwoNumberSumCalculatorSuite) TestShouldCallAndReturnAnArray() {
	arr := []int{1, 2, 3, 4, 5}
	actual := s.TwoNumberSumCalculator.twoNumberSumCalc(arr, 10)
	s.Suite.NotNil(actual)
}
