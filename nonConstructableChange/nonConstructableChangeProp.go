package nonconstructablechange

type NonContructableChange interface {
	NonConstructibleChange(coins []int) int
}

type nonConstructableChange struct {
}

func NewNonContructableChange() *nonConstructableChange {
	return &nonConstructableChange{}
}

func (n *nonConstructableChange) NonConstructibleChange(coins []int) int {
	return 1
}
