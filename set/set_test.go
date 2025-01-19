package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {

	s := New[int]()

	assert.Equal(t, 0, len(s), "Length of a new stack should be 0!")

	s.Add(1)

	assert.Equal(t, 1, len(s), "Length should be 1.")
	assert.Equal(t, true, s.Contains(1), "The set should contain 1.")
	assert.Equal(t, false, s.Contains(2), "The set should not contain 2.")

	s.Add(2)
	s.Remove(1)

	assert.Equal(t, 1, len(s), "Length should be 1 after push.")
	assert.Equal(t, true, s.Contains(2), "The set should contain 2.")
	assert.Equal(t, false, s.Contains(1), "The set should not contain 1.")
}

func TestSetFromIter(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 2}

	s := FromIter(numbers...)

	assert.Equal(t, 4, len(s), "Length should be 4.")
	assert.Equal(t, true, s.Contains(1), "The set should contain 1.")
	assert.Equal(t, true, s.Contains(2), "The set should contain 2.")
	assert.Equal(t, true, s.Contains(3), "The set should contain 3.")
	assert.Equal(t, true, s.Contains(4), "The set should contain 4.")

}
