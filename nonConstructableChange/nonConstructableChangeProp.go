package nonconstructablechange

import "sort"

type NonContructableChange interface {
	NonConstructibleChange(coins []int) int
}

type nonConstructableChange struct {
}

func NewNonContructableChange() *nonConstructableChange {
	return &nonConstructableChange{}
}

func (n *nonConstructableChange) NonConstructibleChange(coins []int) int {
	sort.Ints(coins)
	change := 0
	for _, value := range coins {
		if value > (change + 1) {
			return change + 1
		}
		change += value
	}
	return change + 1
}
