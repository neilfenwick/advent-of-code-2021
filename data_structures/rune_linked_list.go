package data

type RuneLinkedList struct {
	Head *RuneLinkedListNode
	Tail *RuneLinkedListNode
}

type RuneLinkedListNode struct {
	Value rune
	Next  *RuneLinkedListNode
}

func NewRuneLinkedList(initialValues []rune) *RuneLinkedList {
	var (
		result   RuneLinkedList
		previous *RuneLinkedListNode
	)
	for _, char := range initialValues {
		if result.Head == nil {
			result.Head = &RuneLinkedListNode{Value: char}
			previous = result.Head
			continue
		}
		previous.Next = &RuneLinkedListNode{Value: char}
		result.Tail = previous.Next
		previous = previous.Next
	}
	return &result
}

func (linkedList *RuneLinkedList) AppendValue(value rune) {
	if linkedList.Head == nil {
		linkedList.Head = &RuneLinkedListNode{Value: value}
		linkedList.Tail = linkedList.Head
		return
	}
	linkedList.Tail.Next = &RuneLinkedListNode{Value: value}
	linkedList.Tail = linkedList.Tail.Next
}

func (linkedList *RuneLinkedList) String() string {
	result := make([]rune, 0)
	current := linkedList.Head
	for {
		result = append(result, current.Value)
		if current.Next == nil {
			break
		}
		current = current.Next
	}
	return string(result)
}
