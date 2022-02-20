/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if q == nil && p == nil{
        return true
    } else if q == nil || p == nil{
        return false
    }
    result := true
    
    if p.Val != q.Val {
        result = false
    }
    
    if !isSameTree(p.Left, q.Left) {
        result = false
    }
    if !isSameTree(p.Right, q.Right) {
        result = false
    }
    return result
}
