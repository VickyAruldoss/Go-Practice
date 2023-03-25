package threenumbersum

import "sort"

type ThreeNumberSumCalculator interface {
	sum(arry []int, target int) [][]int
}

type threeNumberSumCalculator struct {
}

func NewThreeNumberSumCalculator() *threeNumberSumCalculator {
	return &threeNumberSumCalculator{}
}

func (t *threeNumberSumCalculator) sum(array []int, target int) [][]int {
	sort.Ints(array)
	triplets := [][]int{}
	for i, currentNum := range array {

		leftPointer := i + 1
		rightPointer := len(array) - 1

		for leftPointer < rightPointer {
			sum := currentNum + array[leftPointer] + array[rightPointer]
			if sum == target {
				triplets = append(triplets, []int{currentNum, array[leftPointer], array[rightPointer]})
				leftPointer += 1
				rightPointer -= 1
			} else if sum < target {
				leftPointer += 1
			} else if sum > target {
				rightPointer -= 1
			}
		}

	}
	return triplets
}
