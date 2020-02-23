package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	// 더블 링크드 리스트 예제
	// listExample()

	// 스택 예제
	stackExample()
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
}
