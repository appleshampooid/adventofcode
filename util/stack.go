package util

import (
	"errors"
	"fmt"
)

type Stack[T interface{}] struct {
	data []T
}

func (s *Stack[T]) Push(b T) {
	s.data = append(s.data, b)
}

func (s *Stack[T]) PushN(b []T) {
	s.data = append(s.data, b...)
}

func (s *Stack[T]) Pop() (T, error) {
	length := len(s.data)
	var b T
	if length == 0 {
		return b, errors.New("Stack is empty!")
	}
	b = s.data[length-1]
	s.data = s.data[0 : length-1]
	return b, nil
}

func (s *Stack[T]) PopN(n int) ([]T, error) {
	length := len(s.data)
	var b []T
	if length < n {
		return b, fmt.Errorf("Stack doesn't have %d elements!", n)
	}
	b = s.data[length-n:]
	s.data = s.data[0 : length-n]
	return b, nil
}

func (s *Stack[T]) Peek() (T, error) {
	length := len(s.data)
	var b T
	if length == 0 {
		return b, errors.New("Stack is empty!")
	}
	b = s.data[length-1]
	return b, nil
}

func NewStack[T interface{}]() Stack[T] {
	return Stack[T]{data: make([]T, 0, 5)}
}

func (s Stack[T]) Print() {
	fmt.Println("bottom up, blah")
	fmt.Print("[")
	for _, b := range s.data {
		fmt.Printf("%v\n", b)
	}
	fmt.Print("]\n")
}

func (s Stack[byte]) PrintByteStack() {
	fmt.Println("bottom up, blah")
	fmt.Print("[")
	// for _, b := range s.data {
	fmt.Printf("%s", s.data)
	// }
	fmt.Print("]\n")
}

// func NewStackWithData(data []byte) stack {
// 	return stack{Data: data}
// }
