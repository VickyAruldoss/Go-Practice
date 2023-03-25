package fournumbersum

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestFourNumberSumCalculatorSuite struct {
	suite.Suite
	FourNumberSumCalculator
}

func TestMain(t *testing.T) {
	suite.Run(t, new(TestFourNumberSumCalculatorSuite))
}

func (s *TestFourNumberSumCalculatorSuite) SetupTest() {
	s.FourNumberSumCalculator = NewFourNumberSumCalculator()
}

func (s *TestFourNumberSumCalculatorSuite) TestShouldCallAndReturnAnNotNilValue() {
	array := []int{1, 2, 3, 4, 5, 6, 6}
	actual := s.FourNumberSumCalculator.Calculate(array, 10)
	s.Suite.NotNil(actual)
}
