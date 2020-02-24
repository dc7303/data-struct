package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	// 더블 링크드 리스트 예제
	// listExample()

	// 스택 예제
	// stackExample()

	// 큐 예제
	// queueExample()

	// 트리 예제
	treeExample()
}

func listExample() {
	list := dataStruct.LinkedList{}

	for i := 0; i < 10; i++ {
		list.Append(i)
	}

	_, node := list.Get(4)
	list.Remove(node)
	fmt.Println(list.ToString())

	_, node = list.Get(0)
	list.Remove(node)
	fmt.Println(list.ToString())

	_, node = list.Get(7)
	list.Remove(node)
	fmt.Println(list.ToString())
}

func stackExample() {
	stack := dataStruct.Stack{}
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.ToString())
	val, node := stack.Pop()
	fmt.Println(val, node.Value())
	fmt.Println(stack.ToString())
	val, node = stack.Pop()
	fmt.Println(stack.ToString())
}

func queueExample() {
	queue := dataStruct.Queue{}
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.ToString())
	val, node := queue.Pop()
	fmt.Println(val, node.Value())
	fmt.Println(queue.ToString())
	val, node = queue.Pop()
	fmt.Println(queue.ToString())
}

func treeExample() {
	tree := dataStruct.Tree{}
	initTree(&tree)
	tree.DeepFirstSearch(true)
	fmt.Println()
	tree.DeepFirstSearch(false)
}

/*
		   1
		/  \ \
	   2   3  4
      /\  /\  /\
	 5 6 7 8 9 10
*/
func initTree(t *dataStruct.Tree) {
	val := 1
	t.InitRoot(val)
	val++

	root := t.GetRoot()
	for i := 0; i < 3; i++ {
		root.AppendChild(val)
		val++
	}

	for i := 0; i < len(root.GetChilds()); i++ {
		for j := 0; j < 2; j++ {
			root.GetChilds()[i].AppendChild(val)
			val++
		}
	}
}
