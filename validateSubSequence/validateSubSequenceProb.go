package validatesubsequence

type ValidateSubSequenceProb interface {
	IsValidSubsequence(array []int, sequence []int) bool
}

type validateSubSequence struct {
}

func NewValidateSubSequence() *validateSubSequence {
	return &validateSubSequence{}
}

// O(n) time | O(1) space
func (v *validateSubSequence) IsValidSubsequence(array []int, sequence []int) bool {

	seqIndex := 0
	for _, currentValue := range array {
		if seqIndex == len(sequence) {
			return true
		}
		if sequence[seqIndex] == currentValue {
			seqIndex += 1
		}
	}
	return seqIndex == len(sequence)
}
