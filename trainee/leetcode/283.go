package leetcode

//Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
//Note that you must do this in-place without making a copy of the array.
//
//

//func moveZeroes(nums []int) {
//	ptr1, ptr2 := 0, 0
//	for ptr1 < len(nums) && ptr2 < len(nums) {
//		if nums[ptr1] == 0 && nums[ptr2] != 0 && ptr1 < ptr2 {
//			swap(&nums[ptr1], &nums[ptr2])
//			ptr2++
//			ptr1++
//		} else if nums[ptr1] != 0 {
//			ptr1++
//		} else {
//			ptr2++
//		}
//	}
//	fmt.Println(nums)
//}
//func swap(a, b *int) {
//	temp := *a
//	*a = *b
//	*b = temp
//}

//func moveZeroes(nums []int)  {
//    var slow int
//    for fast := 0; fast < len(nums); fast++ {
//        if nums[fast] != 0 {
//            nums[slow], nums[fast] = nums[fast], nums[slow]
//            slow++
//        }
//    }
//}
