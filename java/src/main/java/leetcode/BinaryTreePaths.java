package leetcode;

import java.util.ArrayList;
import java.util.List;

/**
 * Source  https://leetcode.com/problems/binary-tree-paths/
 * Author  cytian
 * Created 2016-07-28
 *
 * Given a binary tree, return all root-to-leaf paths.
 *
 * For example, given the following binary tree:
 *
 *    1
 *  /   \
 * 2     3
 *  \
 *   5
 * All root-to-leaf paths are:
 * ["1->2->5", "1->3"]
 */

// Definition for a binary tree node.
class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }
}

public class BinaryTreePaths {

    private List<String> paths = new ArrayList<>();

    private void getPaths(TreeNode node, String path) {
        path += node.val;
        if (node.left == null && node.right == null)
            paths.add(path);
        else {
            if (node.left != null)
                getPaths(node.left, path + "->");
            if (node.right != null)
                getPaths(node.right, path + "->");
        }
    }

    public List<String> binaryTreePaths(TreeNode root) {
        if (root != null)
            getPaths(root, "");
        return paths;
    }

}
