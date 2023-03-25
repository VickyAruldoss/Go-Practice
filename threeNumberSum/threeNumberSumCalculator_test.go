package threenumbersum

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestThreeNumberSumCalculatorSuite struct {
	suite.Suite
	ThreeNumberSumCalculator
}

func TestMain(t *testing.T) {
	suite.Run(t, new(TestThreeNumberSumCalculatorSuite))
}

func (suite *TestThreeNumberSumCalculatorSuite) SetupTest() {
	suite.ThreeNumberSumCalculator = NewThreeNumberSumCalculator()
}

func (t *TestThreeNumberSumCalculatorSuite) TestShouldCallAndReturnNonNilObject() {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	expected := t.ThreeNumberSumCalculator.sum(array, 0)
	t.Suite.NotNil(expected)
}

func (t *TestThreeNumberSumCalculatorSuite) TestShouldCallAndReturnNonNilthreeArrays() {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	actual := [][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 1, 5}}
	expected := t.ThreeNumberSumCalculator.sum(array, 0)
	t.Suite.Equal(expected, actual)
}

func (t *TestThreeNumberSumCalculatorSuite) TestShouldCallAndReturnNonNilthree() {
	array := []int{1, 2, 3}
	actual := [][]int{{1, 2, 3}}
	expected := t.ThreeNumberSumCalculator.sum(array, 6)
	t.Suite.Equal(expected, actual)
}
