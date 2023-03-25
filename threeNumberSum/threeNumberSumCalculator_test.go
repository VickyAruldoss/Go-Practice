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
