/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import ("fmt")



type Pair struct {
    Node *TreeNode
    Count int
}


func kthSmallestHelper(root *TreeNode, k int, count int) *Pair {
    
    if root.Left != nil {
        value := kthSmallestHelper(root.Left, k, count)
        if value.Node != nil {
            return value
        }
        count = value.Count
    }
    
    count ++
    if(count == k) {
        return &Pair{Node:root, Count:count}
    }
    
    if root.Right != nil {
        value2 := kthSmallestHelper(root.Right, k, count)
        if value2.Node != nil {
            return value2
        }
        count = value2.Count
    }
    return &Pair{Node:nil, Count:count}
}

func kthSmallest(root *TreeNode, k int) int {
    pair := kthSmallestHelper(root, k, 0)
    return pair.Node.Val
}
