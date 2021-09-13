package main

type stack struct {
	 items []int
}

func (s *stack) push(val int) {
	s.items = append(s.items, val)
}
