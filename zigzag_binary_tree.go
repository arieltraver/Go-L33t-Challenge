/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import ("container/list")

//strategy: use BFS. turn the nodes into a list of lists, then use that to fill the 2d array
//memory: N
//speed: N

func zigzagLevelOrder(root *TreeNode) [][]int {
    
    if (root == nil) {return make([][]int, 0, 0)}
    
    layers := list.New() //list of layers, which are lists.
    
    l1 := list.New()
    l2 := list.New()
    
    layers.PushFront(l1)
    layers.PushFront(l2)
    
    currentLayer := layers.Front() //layer 0, containing the root
    nextLayer := layers.Back() //layer 1, empty as for now.
    
    currentLayer.Value.(*list.List).PushFront(root) //add root to current layer
    current := currentLayer.Value.(*list.List).Front() //point to element containing root.
    
    
    for (current != nil) {
        
        node := current.Value.(*TreeNode)
        nextLayer = currentLayer.Next() //element containing next layer, add children here
        
        if node.Left != nil { nextLayer.Value.(*list.List).PushBack(node.Left)}
        if node.Right != nil {nextLayer.Value.(*list.List).PushBack(node.Right)}
        
        if current.Next() == nil { //end of current layer.
            currentLayer = currentLayer.Next() //onto next layer.
            current = currentLayer.Value.(*list.List).Front() //point to first item in the list.
            lNext := list.New()
            layers.PushBack(lNext) //add an empty list.
        } else {
            current = current.Next() //there are still nodes in the current layer
        }
    }
    
    result := make([][]int, 0, layers.Len()) //avoid re-allocation by setting capacity
    forwards := true
    
    for lyr := layers.Front(); lyr.Value.(*list.List).Front() != nil; lyr = lyr.Next() {
        resultLayer := make([]int, 0, lyr.Value.(*list.List).Len()) //avoid re-allocation by setting capacity.
        if (forwards) { //build the layer forwards
            for e := lyr.Value.(*list.List).Front(); e != nil; e = e.Next() {
                resultLayer = append(resultLayer, e.Value.(*TreeNode).Val)
            }
        } else {
            for e := lyr.Value.(*list.List).Back(); e != nil; e = e.Prev() {
                resultLayer = append(resultLayer, e.Value.(*TreeNode).Val)
            }
        }
        result = append(result, resultLayer)
        forwards = !forwards
    }
    
    return result
    
}
