import ("math");

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minimum(a int, b int ) int {
    if a > b {
        return b;
    } else {
        return a;
    } 
}

func maximum(a int, b int ) int {
    if a > b {
        return a;
    } else {
        return b;
    }
}

func helper(root *TreeNode, min int, max int) bool {
    if root == nil {
        return true
    } else {
        return root.Val > min && root.Val < max && helper(root.Left, min, minimum(root.Val, max)) && helper(root.Right, maximum(root.Val, min), max);
    }
}
func isValidBST(root *TreeNode) (bool) {
    return helper(root, math.MinInt, math.MaxInt)
}
