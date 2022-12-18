package data

type IntBuffer struct {
	buffer         []int
	WriteCursorPos int
}

func NewIntBuffer(size int) *IntBuffer {
	return &IntBuffer{buffer: make([]int, size)}
}

func (b *IntBuffer) Write(val int) {
	b.buffer[b.WriteCursorPos] = val
	b.WriteCursorPos = (b.WriteCursorPos + 1) % len(b.buffer)
}

func (b *IntBuffer) Read(offset int, count int) []int {
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
