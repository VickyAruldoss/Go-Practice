package threenumbersum

type ThreeNumberSumCalculator interface {
	sum(arry []int, target int) []int
}

type threeNumberSumCalculator struct {
}

func NewThreeNumberSumCalculator() *threeNumberSumCalculator {
	return &threeNumberSumCalculator{}
}

func (t *threeNumberSumCalculator) sum(arry []int, target int) []int {
	return arry
}
