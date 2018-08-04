/*
Source  https://leetcode.com/problems/leaf-similar-trees/description/
Author  Tian
Date    2018/08/04
ID      872 Leaf-Similar Trees

Consider all the leaves of a binary tree.  From left to right order, the values of those leaves form a leaf value sequence.

For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).
Two binary trees are considered leaf-similar if their leaf value sequence is the same.
Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

Note:
Both of the given trees will have between 1 and 100 nodes.

Related Topics
Tree, Depth-first Search
*/

package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	root1LN := make([]int, 0) // pull root1's lastNode
	root2LN := make([]int, 0) // pull root2's lastNode
	dfs(root1, &root1LN)
	dfs(root2, &root2LN)

	if len(root1LN) == len(root2LN) {
		// half on for
		for i, j := 0, len(root1LN)-1; i <= j; i, j = i+1, j-1 {
			// if had difference, return false
			if root1LN[i] != root2LN[i] || root1LN[j] != root2LN[j] {
				return false
			}
		}
		return true
	}
	return false
}

// Solution
// Depth first search
func dfs(node *TreeNode, lastNode *[]int) {
	if node == nil {
		return
	}
	// is lastNode
	if node.Left == nil && node.Right == nil {
		*lastNode = append(*lastNode, node.Val) // pull lastNode to slices
		return
	}
	dfs(node.Left, lastNode)
	dfs(node.Right, lastNode)
}

func main() {

}
