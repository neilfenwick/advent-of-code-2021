package data

type Stack struct {
	stack []interface{}
}

func NewStack() *Stack {
	stack := make([]interface{}, 0, 100)
	return &Stack{stack: stack}
}

func NewStackFromItems(items []interface{}) *Stack {
	stack := NewStack()
	stack.stack = items
	return stack
}

func (s *Stack) Push(item interface{}) {
	s.stack = append(s.stack, item)
}

func (s *Stack) Pop() (value interface{}, found bool) {
	item, found := s.Peek()
	if found {
		s.stack = s.stack[:len(s.stack)-1]
	}
	return item, found
}

func (s *Stack) Peek() (value interface{}, found bool) {
	length := len(s.stack)
	if length == 0 {
		return nil, false
	}

	item := s.stack[length-1]
	return item, true
}

func (s *Stack) Copy() *Stack {
	newCopy := make([]interface{}, len(s.stack))
	copy(newCopy, s.stack)
	result := &Stack{stack: newCopy}
	return result
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) Items() []interface{} {
	return s.stack
}
