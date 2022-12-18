package data

type Queue struct {
	queue []interface{}
}

func NewQueue() *Queue {
	stack := make([]interface{}, 0, 100)
	return &Queue{queue: stack}
}

func (s *Queue) Push(item interface{}) {
	s.queue = append(s.queue, item)
}

func (s *Queue) Pop() (value interface{}, found bool) {
	item, found := s.Peek()
	if found {
		s.queue = s.queue[1:]
	}
	return item, found
}

func (s *Queue) Peek() (value interface{}, found bool) {
	length := len(s.queue)
	if length == 0 {
		return nil, false
	}

	item := s.queue[0]
	return item, true
}

func (s *Queue) Copy() *Queue {
	newCopy := make([]interface{}, len(s.queue))
	copy(newCopy, s.queue)
	result := &Queue{queue: newCopy}
	return result
}

func (s *Queue) Size() int {
	return len(s.queue)
}

func (s *Queue) Items() []interface{} {
	return s.queue
}
