package binary_tree

import (
	"carerti/datastructure"
	"fmt"
)

func DisplayLevelOrderTree(root *TreeNode) {
	if root == nil{
		fmt.Println("Tree is empty")
		return
	}

	result, numberOfNodes := levelOrder2D(root)
	if numberOfNodes > 0 {
		fmt.Println("Tree: ", result)

	}
	fmt.Println("Is tree Balanced: ", IsbBalanced(nil))
}

func levelOrder2D(root *TreeNode) ([][]int, int) {
	var level []int
	var result [][]int
	queue := new(datastructure.Queue)
	queue.Enqueue(root)
	levelLength := queue.Size()
	numberOfNodes := levelLength
	for !queue.IsEmpty() {

		if levelLength == 0 {
			result = append(result, level)
			levelLength = queue.Size()
			numberOfNodes += levelLength
			level = []int{}
		}
		tempNode := queue.Dequeue().(*TreeNode)
		level = append(level, tempNode.Val)
		levelLength--
		if tempNode.Left != nil {
			queue.Enqueue(tempNode.Left)
		}
		if tempNode.Right != nil {
			queue.Enqueue(tempNode.Right)
		}
	}
	result = append(result, level)
	return result, numberOfNodes
}
