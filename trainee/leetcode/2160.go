package leetcode

//You are given a positive integer num consisting of exactly four digits.
//Split num into two new integers new1 and new2 by using the digits found in num.
//Leading zeros are allowed in new1 and new2, and all the digits found in num must be used.
//For example, given num = 2932, you have the following digits: two 2's, one 9 and one 3.
//Some of the possible pairs [new1, new2] are [22, 93], [23, 92], [223, 9] and [2, 329].
//Return the minimum possible sum of new1 and new2.

//func minimumSum(num int) int {
//    m := make([]int, 0, 4)
//    for i := 0; i < 4; i++ {
//        m = append(m, num % 10)
//        num = num/10
//    }
//    sort.Slice(m, func(i, j int) bool{
//        return m[i] < m[j]
//    })
//    return m[0] * 10 + m[2] + m[1] * 10 + m[3]
//}
