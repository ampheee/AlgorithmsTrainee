package leetcode

//Given the root of a binary tree, return the sum of values of its deepest leaves.
//
///**
// * Definition for a binary tree node.
// * type TreeNode struct {
// *     Val int
// *     Left *TreeNode
// *     Right *TreeNode
// * }
// */
//func deepestLeavesSum(root *TreeNode) int {
//    if root == nil {
//        return 0
//    }
//    queue := []*TreeNode{root}
//    sum := 0
//    for len(queue) > 0 {
//        sum = 0
//        nodesCurrent := len(queue)
//        for i := 0; i < nodesCurrent; i++ {
//            curr := queue[0]
//            queue = queue[1:]
//            sum += curr.Val
//            if curr.Left != nil {
//                queue = append(queue, curr.Left)
//            }
//            if curr.Right != nil {
//                queue = append(queue, curr.Right)
//            }
//        }
//    }
//    return sum
//}
