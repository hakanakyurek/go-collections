package collections

import (
	"hakanakyurek/go-collections/set"
	"hakanakyurek/go-collections/stack"
)

type Stack[T any] = stack.Stack[T]
type Set[T comparable] = set.Set[T]

func NewStack[T any]() Stack[T] {
	return stack.New[T]()
}

func NewSet[T comparable]() Set[T] {
	return set.New[T]()
}
