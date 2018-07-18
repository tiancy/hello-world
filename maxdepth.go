/*
Source  https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
Author  Tian
Created 2018/07/18

104. Maximum Depth of Binary Tree

Given a binary tree, find its maximum depth.
The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
Note: A leaf is a node with no children.

Related Topics 
Tree, Depth-first Search
*/
package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var depth int

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth = 0
	dfs(root, 0)
	// dfs2(root, 1)
	return depth
}

// solution 1
func dfs(node *TreeNode, level int) {
	if node == nil {
		if depth < level {
			depth = level
		}
		return
	}

	level++
	dfs(node.Left, level)
	dfs(node.Right, level)
}

// solution 2
func dfs2(node *TreeNode, level int) {
	if node.Left == nil && node.Right == nil {
		if depth < level {
			depth = level
		}
	} else {
		if node.Left != nil {
			dfs(node.Left, level + 1)
		}
		if node.Right != nil {
			dfs(node.Right, level + 1)
		}
	}
}

func main() {
	n2 := &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}
	n5 := &TreeNode{5, &TreeNode{6, nil, nil}, n2}
	n0 := &TreeNode{Val:0}
	root := TreeNode{3, n5, &TreeNode{1, n0, &TreeNode{8, nil, nil}}}
	d := maxDepth(&root)
	fmt.Println(d)
}