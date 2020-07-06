package main

import (
	"carerti/binary_tree"
)

func main() {
	root := newNode(1)
	root.Left = newNode(2)
	root.Right = newNode(3)
	root.Left.Left = newNode(4);
	root.Left.Right = newNode(5);
	root.Right.Left = newNode(6);
	root.Right.Right = newNode(7);
	root.Left.Left.Left = newNode(8);
	root.Left.Left.Right = newNode(9);
	root.Left.Right.Left = newNode(10);
	root.Left.Right.Right = newNode(11);
	root.Right.Left.Left = newNode(12);
	root.Right.Left.Right = newNode(13)
	root.Right.Right.Left = newNode(14)
	root.Right.Right.Right = newNode(15)
	root.Right.Right.Right.Right = newNode(16)
	// binary_tree.DisplayLevelOrderTree(nil) // when tree is empty
	binary_tree.DisplayLevelOrderTree(root)
}

func newNode(data int) *binary_tree.TreeNode {
	node := new(binary_tree.TreeNode)
	node.Val = data
	node.Left = nil
	node.Right = nil
	return node
}
