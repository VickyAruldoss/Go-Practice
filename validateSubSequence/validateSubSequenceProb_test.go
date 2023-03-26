package validatesubsequence

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestValidateSubSequenceSuite struct {
	suite.Suite
	ValidateSubSequenceProb
}

func TestMain(t *testing.T) {
	suite.Run(t, new(TestValidateSubSequenceSuite))
}

func (s *TestValidateSubSequenceSuite) SetupTest() {
	s.ValidateSubSequenceProb = NewValidateSubSequence()
}

func (s *TestValidateSubSequenceSuite) Test_ShouldCallAndReturnBoolValue() {
	expected := s.ValidateSubSequenceProb.IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10})
	s.Suite.True(expected)
}
