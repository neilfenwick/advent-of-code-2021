package data

type CircularBuffer struct {
	buffer         []interface{}
	WriteCursorPos int
}

func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{buffer: make([]interface{}, size)}
}

func (b *CircularBuffer) Write(val interface{}) {
	b.buffer[b.WriteCursorPos] = val
	b.WriteCursorPos = (b.WriteCursorPos + 1) % len(b.buffer)
}

func (b *CircularBuffer) Read(offset int, count int) []interface{} {
	startPos := (b.WriteCursorPos + offset) % len(b.buffer)
	if startPos < 0 {
		startPos = len(b.buffer) + startPos
	}

	if startPos+count <= len(b.buffer) {
		return b.buffer[startPos : startPos+count]
	}

	tail := b.buffer[startPos:]
	wrap := b.buffer[0 : count-len(tail)]
	tail = append(tail, wrap...)
	return tail
}

func (b *CircularBuffer) Size() int {
	return len(b.buffer)
}
