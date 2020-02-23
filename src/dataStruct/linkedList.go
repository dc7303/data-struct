package dataStruct

import (
	"strconv"
)

type Node struct {
	next *Node
	pre  *Node
	val  int
}

func (n *Node) Value() int {
	return n.val
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) Append(val int) {
	// 최초 생성일 때
	newNode := &Node{val: val}
	if l.root == nil {
		l.root = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	newNode.pre = l.tail
	l.tail = newNode
}

func (l *LinkedList) Remove(target *Node) {
	// pre가 없을 때
	// for node.next != target {
	// 	node = node.next
	// }
	if l.root == target {
		l.root = l.root.next
		l.root.pre = nil
		return
	}

	if l.tail == target {
		pre := l.tail.pre
		pre.next = nil
		l.tail = target.pre
		return
	}

	pre := target.pre
	next := target.next
	pre.next, next.pre = next, pre
	target.next, target.pre = nil, nil
}

func (l *LinkedList) Get(index int) (int, *Node) {
	node := l.root
	if index == 0 {
		return node.val, node
	}

	for i := 0; i < index; i++ {
		node = node.next
	}

	if node == nil {
		return 0, nil
	}

	return node.val, node
}

func (l *LinkedList) ToString() string {
	result := "["

	node := l.root
	if node == nil {
		return result + "]"
	}

	for node != nil {
		result += (strconv.Itoa(node.val) + " ")
		node = node.next
	}

	result = result[:len(result)-1] + "]"
	return result
}

func (l *LinkedList) GetRootNode() *Node {
	return l.root
}

func (l *LinkedList) GetTailNode() *Node {
	return l.tail
}
