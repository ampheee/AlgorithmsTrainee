package leetcode

//There are n kids with candies. You are given an integer array candies,
//where each candies[i] represents the number of candies the ith kid has,
//and an integer extraCandies, denoting the number of extra candies that you have.
//Return a boolean array result of length n, where result[i] is true if,
//after giving the ith kid all the extraCandies,
//they will have the greatest number of candies among all the kids, or false otherwise.
//Note that multiple kids can have the greatest number of candies.

//func kidsWithCandies(candies []int, extraCandies int) []bool {
//    var max = candies[0]
//    for _, val := range candies {
//        if val > max {
//            max = val
//        }
//    }
//    res := make([]bool, 0, len(candies))
//    for _, v := range candies {
//        if v +  extraCandies >= max {
//            res = append(res, true)
//        } else {
//            res = append(res, false)
//        }
//    }
//    return res
//}
