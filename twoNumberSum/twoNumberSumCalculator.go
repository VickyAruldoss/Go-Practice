package twonumbersum

import "sort"

type TwoNumberSumCalculator interface {
	twoNumberSumCalc(array []int, target int) []int
	twoNumberSumCalcSol2(array []int, target int) []int
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

func (t *twoNumberSumCalculator) twoNumberSumCalcSol2(array []int, target int) []int {
	sort.Ints(array)
	leftPointer := 0
	rightPointer := len(array) - 1

	for leftPointer != rightPointer {
		sum := array[leftPointer] + array[rightPointer]
		if sum == target {
			return []int{array[rightPointer], array[leftPointer]}
		} else if sum < target {
			leftPointer += 1
		} else if sum > target {
			rightPointer -= 1
		}
	}
	return []int{}

}
