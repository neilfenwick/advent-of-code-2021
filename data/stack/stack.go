package stack

type Stack struct {
	stack []interface{}
}

func NewStack() *Stack {
	stack := make([]interface{}, 0, 100)
	return &Stack{stack: stack}
}

func (s *Stack) Push(item interface{}) {
	s.stack = append(s.stack, item)
}

func (s *Stack) Pop() (interface{}, bool) {
	item, found := s.Peek()
	if found {
		s.stack = s.stack[:len(s.stack)-1]
	}
	return item, found
}

func (s *Stack) Peek() (interface{}, bool) {
	length := len(s.stack)
	if length == 0 {
		return nil, false
	}

	item := s.stack[length-1]
	return item, true
}
