package main

import "strings"

func simplifyPath(path string) string {
	stack := New[string]()
	els := strings.Split(path, "/")
	for i := range len(els) {
		if els[i] == ".." {
			if !stack.IsEmpty() {
				_ = stack.Pop()
			}
			continue
		}
		if els[i] != "" && els[i] != "." {
			stack.Push(els[i])
		}
	}
	return "/" + strings.Join(stack.GetData(), "/")

}

type Stack[T any] struct {
	data []T
}

// New создает новый стек
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push добавляет элемент в стек
func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop извлекает верхний элемент (предполагается, что стек не пустой)
func (s *Stack[T]) Pop() T {
	index := len(s.data) - 1
	value := s.data[index]
	s.data = s.data[:index]
	return value
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) GetData() []T {
	return s.data
}
