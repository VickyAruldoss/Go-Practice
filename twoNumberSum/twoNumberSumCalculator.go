package twonumbersum

type TwoNumberSumCalculator interface {
	twoNumberSumCalc(array []int, target int) []int
}

type twoNumberSumCalculator struct {
}

func NewTwoNumberSumCalculator() *twoNumberSumCalculator {
	return &twoNumberSumCalculator{}
}

func (t *twoNumberSumCalculator) twoNumberSumCalc(array []int, target int) []int {
	return array
}
