package leetcode

//Given a 0-indexed integer array nums, find a 0-indexed integer array answer where:
//answer.length == nums.length.
//answer[i] = |leftSum[i] - rightSum[i]|.
//Where:
//leftSum[i] is the sum of elements to the left of the index i in the array nums. If there is no such element, leftSum[i] = 0.
//rightSum[i] is the sum of elements to the right of the index i in the array nums. If there is no such element, rightSum[i] = 0.
//Return the array answer.

//func leftRigthDifference(nums []int) []int {
//    l := len(nums)
//    res := make([]int, 0, l)
//    rightSum := make([]int, 0, l)
//    leftSum := make([]int, 0, l)
//    leftSum = append(leftSum, 0)
//    rightSum = append(rightSum, 0)
//    for i := 1; i < l; i++ {
//        leftSum = append(leftSum, leftSum[i - 1] + nums[i - 1])
//        rightSum = append(rightSum, rightSum[i - 1] + nums[l - i])
//    }
//    for i := 0; i < l; i++ {
//        res = append(res, abs(rightSum[l - i - 1] - leftSum[i]))
//    }
//    return res
//}
//
//func abs(a int) int {
//    if a < 0 {
//        return a * -1
//    }
//    return a
//}
