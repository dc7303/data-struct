package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	listExample()
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
