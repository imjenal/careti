package binary_tree

import (
	"math"
)
func IsbBalanced(root *TreeNode) bool {
	_, ok := getDepthAndBalance(root)
	return ok
}

func getDepthAndBalance(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftHeight, isLeftBalanced := getDepthAndBalance(root.Left)
	rightHeight, isRightBalanced := getDepthAndBalance(root.Right)

	if !isLeftBalanced || !isRightBalanced { //when left or right aren't balanced, then tree isn't balanced. return false
		return 0, false
	}

	if absoluteDifference(leftHeight, rightHeight) > 1 {
		return 0, false
	}
	return findMaximumOf(leftHeight+1, rightHeight+1), true
}

func absoluteDifference(leftHeight, rightHeight int) int {
	return int(math.Abs(float64(leftHeight) - float64(rightHeight)))
}

func findMaximumOf(leftHeight, rightHeight int) int {
	if leftHeight > rightHeight {
		return leftHeight
	}
	return rightHeight
}
