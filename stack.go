package main

type stack struct {
	 items []int
}

func (s *stack) push(val int) {
	s.items = append(s.items, val)
}

func (s *stack) pop() int {
	l := len(s.items)-1
	remove := s.items[l]
	s.items = s.items[:l]
	return  remove
}

func (s *stack) isEmpty() bool {

}