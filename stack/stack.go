package stack

import (
	"errors"
)

type Stack[T any] []T
var ErrEmptyStack = errors.New("Stack is empty!")

func (s *Stack[T]) Push(element T) {
	*s = append(*s,element)
}

func (s *Stack[T]) Pop() (T,error) {
	if len(*s) == 0 {
		var zero T
		return zero,ErrEmptyStack
	}
	element := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return element,nil
}

func (s *Stack[T]) Peek() (T,error) {
	if len(*s) == 0 {
		var zero T
		return zero,ErrEmptyStack
	}
	return (*s)[len(*s)-1],nil
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func New[T any](slice []T) Stack[T] {
	stack := Stack[T]{}
	for _,element := range slice {
		stack.Push(element)
	}
	return stack
}