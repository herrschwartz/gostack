package stack

type stack struct {
	length int
	st []  int
}

func (s *stack) Push(v int) {
	s.length++
	s.st = append(s.st, v)
}

func (s *stack) Pop() int {
	l := s.length
	out := s.st[l-1]
	s.st = s.st[:l-1]
	s.length--
	return out
}

func (s *stack)isEmpty() bool{
	return s.length==0
}

func newStack() *stack {
	s := new(stack)
	s.st = make([]int, 0)
	return s
}

