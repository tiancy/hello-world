/*
Source  https://leetcode.com/problems/find-bottom-left-tree-value/description/
Author  Tian
Date    2018/09/08

513. Find Bottom Left Tree Value

Given a binary tree, find the leftmost value in the last row of the tree.

Example 1:
Input:

    2
   / \
  1   3

Output: 1

Example 2:
Input:

        1
       / \
      2   3
     /   / \
    4   5   6
       /
      7

Output: 7

Related Topics
Tree, Depth-first Search, Breadth-first Search
*/

package bottomleftvalue

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var bl, v int

func findBottomLeftValue(root *TreeNode) int {
	bl, v = 0, 0
	dfs(root, 0)

	return v
}

func dfs(node *TreeNode, l int) {
	if node == nil {
		return
	}

	l++ // level
	dfs(node.Left, l)
	// record bottom left, and val
	if l > bl {
		bl = l
		v = node.Val
	}

	dfs(node.Right, l)
}
