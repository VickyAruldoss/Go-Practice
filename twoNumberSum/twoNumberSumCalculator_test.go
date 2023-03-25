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

func (s *TwoNumberSumCalculatorSuite) TestShouldReturnSumOfTwoNumbers() {
	arr := []int{3, 5, -4, 8, 11, 1, -1, 6}
	expected := []int{11, -1}
	actutal := s.TwoNumberSumCalculator.twoNumberSumCalc(arr, 10)
	s.Suite.Equal(expected, actutal)
}

func (s *TwoNumberSumCalculatorSuite) TestShouldAddSumOfTwoNumbersWhenLengthIsTwo() {
	arr := []int{4, 6}
	expected := []int{4, 6}
	actutal := s.TwoNumberSumCalculator.twoNumberSumCalc(arr, 10)
	s.Suite.Equal(expected, actutal)
}

func (s *TwoNumberSumCalculatorSuite) TestShouldAddSumOfTwoNumbersWhenLengthIsHuge() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{8, 9}
	actutal := s.TwoNumberSumCalculator.twoNumberSumCalc(arr, 17)
	s.Suite.Equal(expected, actutal)
}

func (s *TwoNumberSumCalculatorSuite) TestShouldSumOfTwoNumbersForSolutionTwo() {
	arr := []int{3, 5, -4, 8, 11, 1, -1, 6}
	expected := []int{11, -1}
	actutal := s.TwoNumberSumCalculator.twoNumberSumCalcSol2(arr, 10)
	s.Suite.Equal(expected, actutal)
}

func (s *TwoNumberSumCalculatorSuite) TestShouldAddSumOfTwoNumbersWhenLengthIsTwoSolTwo() {
	arr := []int{4, 6}
	expected := []int{6, 4}
	actutal := s.TwoNumberSumCalculator.twoNumberSumCalcSol2(arr, 10)
	s.Suite.Equal(expected, actutal)
}

func (s *TwoNumberSumCalculatorSuite) TestShouldAddSumOfTwoNumbersWhenLengthIsHugeSolTwo() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{9, 8}
	actutal := s.TwoNumberSumCalculator.twoNumberSumCalcSol2(arr, 17)
	s.Suite.Equal(expected, actutal)
}
