package fournumbersum

type FourNumberSumCalculator interface {
	Calculate(array []int, target int) [][]int
}

type fourNumberSumCalculator struct {
}

func NewFourNumberSumCalculator() *fourNumberSumCalculator {
	return &fourNumberSumCalculator{}
}

func (f *fourNumberSumCalculator) Calculate(array []int, target int) [][]int {
	return [][]int{}
}
