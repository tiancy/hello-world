/*
Source  https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/description/
Author  tian
Created 2018/07/15

863. All Nodes Distance K in Binary Tree

We are given a binary tree (with root node root), a target node, and an integer value K.
Return a list of the values of all nodes that have a distance K from the target node.  The answer can be returned in any order.

Example 1:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
Output: [7,4,1]

Related Topics
Tree, Depth-first Search, Breadth-first Search
*/

package main

import (
	"fmt"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var parents map[TreeNode]TreeNode

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if root == nil || target == nil {
		return nil
	}

	parents = make(map[TreeNode]TreeNode)
	dfs(root, &TreeNode{Val: -1})

	queue := make([]TreeNode, 0)
	queue = append(queue, TreeNode{Val: -1}) // push null
	queue = append(queue, *target)

	seen := make(map[TreeNode]TreeNode)
	seen[*target] = *target

	dist := 0
	for len(queue) != 0 {
		node := queue[0]  // Top (just get next element, don't remove it
		queue = queue[1:] // Discard top element
		if node.Val == -1 {
			if dist == K {
				ans := make([]int, 0)
				for _, v := range queue {
					ans = append(ans, v.Val)
				}
				return ans
			}
			queue = append(queue, TreeNode{Val: -1}) // push null, because add one depth after to ==
			dist++
		} else {
			// append node.Left, node.Right, node Parent to queue
			// if have one is distanceK when all is distanceK
			if node.Left != nil {
				// v != *node.Leftp, v is new object
				// if v, r := seen[*node.Left]; !r {
				// 	queue = append(queue, *node.Left)
				// 	seen[v] = v
				// }
				if _, r := seen[*node.Left]; !r {
					queue = append(queue, *node.Left)
					seen[*node.Left] = *node.Left
				}
			}
			if node.Right != nil {
				if _, r := seen[*node.Right]; !r {
					queue = append(queue, *node.Right)
					seen[*node.Right] = *node.Right
				}
			}
			p := parents[node]
			// node is't root
			if p.Val != -1 {
				if _, r := seen[p]; !r {
					queue = append(queue, p)
					seen[p] = p
				}
			}
		}
	}

	return nil
}

// dfs is depth first search
func dfs(node *TreeNode, parent *TreeNode) {
	if node != nil {
		parents[*node] = *parent
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
}

func main() {
	n2 := &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}
	n5 := &TreeNode{5, &TreeNode{6, nil, nil}, n2}
	n0 := &TreeNode{0, nil, nil}
	root := TreeNode{3, n5, &TreeNode{1, n0, &TreeNode{8, nil, nil}}}
	r := distanceK(&root, n0, 4)
	fmt.Println(r)

	// n3 := &TreeNode{3, nil, nil}
	// root := TreeNode{0, &TreeNode{2, nil, nil}, &TreeNode{1, n3, nil}}
	// r := distanceK(&root, n3, 3)
	// fmt.Println(r)
}
