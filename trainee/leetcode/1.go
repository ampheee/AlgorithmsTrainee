package leetcode

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//You can return the answer in any order.
//
//

//func twoSum(nums []int, target int) []int {
//    var res []int
//    m := make(map[int]int)
//    for i := 0; i < len(nums); i++ {
//        check := target - nums[i]
//        if val, ok := m[check]; ok {
//            res = []int{i, val}
//            break
//        }
//        m[nums[i]] = i
//    }
//    fmt.Println(m)
//    return res
//}
