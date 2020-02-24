package dataStruct

type BinaryTreeNode struct {
	left *BinaryTreeNode
	right *BinaryTreeNode
	val int
}

func (bt *BinaryTreeNode) AddNode(val int) {
	if bt.val > val {
		if bt.left == nil {
			bt.left = &BinaryTreeNode{val: val}
		} else {
			bt.left.AddNode(val)
		}
	} else {
		if bt.right == nil {
			bt.right = &BinaryTreeNode{val: val}
		} else {
			bt.right.AddNode(val)
		}
	}
}

func (bt *BinaryTreeNode) Search(val int) bool {
	if bt.val == val {
		return true
	} else if bt.val < val {
		if bt.right != nil {
			return bt.right.Search(val)
		}
		return false
	} else if bt.val > val {
		if bt.left != nil {
			return bt.left.Search(val)
		}
		return false
	}
	return false
}

type BinaryTree struct {
	root *BinaryTreeNode
}

func (b *BinaryTree) GetRoot() *BinaryTreeNode {
	return b.root
}

func (b *BinaryTree) InitRoot(val int) {
	if b.root == nil {
		b.root = &BinaryTreeNode{val: val}
	}
}

/*
O(log2N)
*/
func (b *BinaryTree) BinarySearchTree(val int) bool {
	return b.root.Search(val)
}
