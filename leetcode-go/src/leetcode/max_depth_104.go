package leetcode

import (
	"container/list"
	"./common"
)

//Given a binary tree, find its maximum depth.
//The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//Given binary tree [3,9,20,null,null,15,7],
//给定一个二叉树，找出其最大深度。
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数
//      3
//     / \
//    9  20
//       /  \
//      15   7
//  return its depth = 3.
// 递归
func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	} else {
		leftHeight := maxDepth(root.Left)
		rightHeight := maxDepth(root.Right)
		return max(leftHeight, rightHeight) + 1
	}
}

// 迭代
func maxDepth2(root *common.TreeNode) int {
	stack := list.New()
	if root != nil {
		stack.PushBack(&Pair{
			First: root,
			Second: 1,
		})
	}

	depth := 0
	for stack != nil && stack.Len() != 0 {
		currentElement := stack.Front()
		stack.Remove(currentElement)
		root = currentElement.Value.(*Pair).First
		currentDepth := currentElement.Value.(*Pair).Second
		if root != nil {
			depth = max(currentDepth, depth)
			stack.PushBack(&Pair{
				First: root.Left,
				Second: currentDepth + 1,
			})
			stack.PushBack(&Pair{
				First: root.Right,
				Second: currentDepth + 1,
			})
		}
	}
	return depth
}

type Pair struct {
	First *common.TreeNode
	Second int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
