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
	// treeExample()

	// 이진 트리 예제
	// binaryTreeExample()

	// 힙 예제
	// heapExample()

	// 맵 예제
	mapExample()
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
	fmt.Println()
	tree.BreadthFirstSearch()
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

func binaryTreeExample() {
	tree := dataStruct.BinaryTree{}
	tree.InitRoot(5)
	root := tree.GetRoot()
	root.AddNode(3)
	root.AddNode(4)
	root.AddNode(2)
	root.AddNode(8)
	root.AddNode(7)
	root.AddNode(6)
	root.AddNode(10)
	root.AddNode(9)

	findVal := 20
	if tree.BinarySearchTree(findVal) {
		fmt.Printf("Found value = %d", findVal)
	} else {
		fmt.Printf("Not Found value = %d", findVal)
	}
}

func heapExample() {
	h := &dataStruct.Heap{}
	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(4)
	h.Push(5)

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())

	h.PrintList()
}

func mapExample() {
	hashMap := dataStruct.HashMap{}
	hashMap.Put("Margurt", "최동철")
	hashMap.Put("Domingo", "오세용")
	hashMap.Put("NariYoo", "유나리")
	hashMap.Put("NotFound404", "곽대종")

	fmt.Println(hashMap.Get("Margurt"))
	fmt.Println(hashMap.Get("Domingo"))
	fmt.Println(hashMap.Get("NariYoo"))
	fmt.Println(hashMap.Get("NotFound404"))
}
