package sortedsquaredarray

import "sort"

type SortedSquaredArrayProb interface {
	SortedSquaredArray(array []int) []int
}

type sortedSquaredArrayProb struct {
}

func NewSortedSquaredArray() *sortedSquaredArrayProb {
	return &sortedSquaredArrayProb{}
}

func (s *sortedSquaredArrayProb) SortedSquaredArray(array []int) []int {
	result := []int{}
	for _, curentValue := range array {
		result = append(result, curentValue*curentValue)
	}
	sort.Ints(result)
	return result
}
