package leetcode

//Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it.
//That is, for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].

//func smallerNumbersThanCurrent(nums []int) []int {
//    l := len(nums)
//    res := make([]int, 0, l)
//    for i := 0; i < l; i++ {
//        count := 0
//        for j := 0; j < l; j++ {
//            if j != i &&nums[i] > nums[j] {
//                count++
//            }
//        }
//        res = append(res, count)
//    }
//    return res
//}
