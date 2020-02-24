package dataStruct

type Queue struct {
	list LinkedList
}

func (q *Queue) Push(val int) {
	list := &q.list
	list.Append(val)
}

func (q *Queue) Pop() (int, *Node) {
	list := &q.list
	if list.IsEmpty() {
		return 0, nil
	}

	root := list.GetRootNode()
	reVal, reRoot := root.Value(), *root
	list.Remove(root)
	return reVal, &reRoot
}

func (q *Queue) IsEmpty() bool {
	list := &q.list
	return list.IsEmpty()
}

func (q *Queue) ToString() string {
	list := &q.list
	return list.ToString()
}
