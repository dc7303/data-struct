package dataStruct

import "fmt"

type TreeNode struct {
	val    int
	childs []*TreeNode
}

func (tn *TreeNode) Value() int {
	return tn.val
}

func (tn *TreeNode) GetChilds() []*TreeNode {
	return tn.childs
}

func (tn *TreeNode) AppendChild(val int) {
	tn.childs = append(tn.childs, &TreeNode{val: val})
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) InitRoot(val int) {
	if t.root == nil {
		t.root = &TreeNode{val: val}
	}
}

func (t *Tree) GetRoot() *TreeNode {
	return t.root
}

func (t *Tree) DeepFirstSearch(recursive bool) {
	if recursive {
		recursiveDFS(t.root)
	} else {
		stackDFS(t.root)
	}
}

func recursiveDFS(node *TreeNode) {
	fmt.Printf("%d -> ", node.val)

	for i := 0; i < len(node.childs); i++ {
		recursiveDFS(node.childs[i])
	}
}

func stackDFS(node *TreeNode) {
	stack := []*TreeNode{}
	stack = append(stack, node)

	for len(stack) > 0 {
		var last *TreeNode
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]

		fmt.Printf("%d -> ", last.val)

		for j := 0; j < len(last.childs); j++ {
			stack = append(stack, last.childs[j])
		}
	}
}
