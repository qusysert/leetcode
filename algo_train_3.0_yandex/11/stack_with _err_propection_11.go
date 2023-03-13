package main

import "fmt"

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(in int) string {
	s.data = append(s.data, in)
	return "ok"
}

func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("error")
	}
	popped := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return popped, nil
}

func (s *Stack) Back() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("error")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Clear() string {
	s.data = []int{}
	return "ok"
}
