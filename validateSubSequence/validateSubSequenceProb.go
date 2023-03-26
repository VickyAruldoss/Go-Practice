package validatesubsequence

type ValidateSubSequenceProb interface {
	IsValidSubsequence(array []int, sequence []int) bool
}

type validateSubSequence struct {
}

func NewValidateSubSequence() *validateSubSequence {
	return &validateSubSequence{}
}

func (v *validateSubSequence) IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	return false
}
