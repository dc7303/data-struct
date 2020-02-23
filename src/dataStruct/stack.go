package dataStruct

type Stack struct {
	list LinkedList
}

func (s *Stack) Push(val int) {
	list := &s.list
	list.Append(val)
}

func (s *Stack) Pop() (int, Node) {
	tail := s.list.GetTailNode()
	reNode, reVal := *tail, (*tail).val
	s.list.Remove(tail)
	return reVal, reNode
}

func (s *Stack) ToString() string {
	return s.list.ToString()
}
