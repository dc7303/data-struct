package dataStruct

type Stack struct {
	list LinkedList
}

func (s *Stack) Push(val int) {
	list := &s.list
	list.Append(val)
}

func (s *Stack) Pop() (int, *Node) {
	list := &s.list
	if list.IsEmpty() {
		return 0, nil
	}

	tail := list.GetTailNode()
	reVal, reNode := (*tail).val, *tail
	list.Remove(tail)
	return reVal, &reNode
}

func (s *Stack) IsEmpty() bool {
	list := &s.list
	return list.IsEmpty()
}

func (s *Stack) ToString() string {
	list := &s.list
	return list.ToString()
}
