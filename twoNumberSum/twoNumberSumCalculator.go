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
	storedValues := map[int]bool{}
	for _, currentValue := range array {
		potentialValue := target - currentValue
		if _, ok := storedValues[potentialValue]; ok {
			return []int{potentialValue, currentValue}
		}
		storedValues[currentValue] = true
	}
	return []int{}
}
