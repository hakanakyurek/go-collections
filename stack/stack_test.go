package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	s := New[int]()

	assert.Equal(t, 0, len(s), "Length of a new stack should be 0!")
	assert.Equal(t, true, s.Empty(), "Empty function should recognize stack is empty.")

	s.Push(1)

	assert.Equal(t, 1, len(s), "Length should be 1 after push.")
	assert.Equal(t, 1, s.Size(), "Size function should correctly get the size of the stack.")
	assert.Equal(t, s.Top(), 1, "Top element should be peeked as 1.")
	assert.Equal(t, s.Pop(), 1, "Top element should be peeked as 1.")
	assert.Equal(t, 0, len(s), "Length should be 0 after pop.")

	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 3, s.Size(), "Size function should correctly get the size of the stack.")
	assert.Equal(t, s.Pop(), 3, "Top element should be peeked as 3.")
	assert.Equal(t, s.Top(), 2, "Top element should be peeked as 2.")
	assert.Equal(t, 2, s.Size(), "Size function should correctly get the size of the stack.")
	assert.Equal(t, false, s.Empty(), "Empty function should recognize stack is not empty.")

}
