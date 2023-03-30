package nonconstructablechange

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestNonConstructableChangePropSuite struct {
	suite.Suite
	NonContructableChange
}

func TestMain(t *testing.T) {
	suite.Run(t, new(TestNonConstructableChangePropSuite))
}

func (s *TestNonConstructableChangePropSuite) SetupTest() {
	s.NonContructableChange = NewNonContructableChange()
}

func (s *TestNonConstructableChangePropSuite) Test_ShouldAndReturnAnonInterger() {
	actual := s.NonContructableChange.NonConstructibleChange([]int{1, 2, 3, 4, 5})
	s.Suite.NotNil(actual)
}

func (s *TestNonConstructableChangePropSuite) Test_ShouldAndReturnAnonInterger_Retun_20() {
	actual := s.NonContructableChange.NonConstructibleChange([]int{5, 7, 1, 1, 2, 3, 22})
	s.Suite.Equal(20, actual)
}
