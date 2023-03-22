package leetcode

//You're given strings jewels representing the types of stones that are jewels,
//and stones representing the stones you have. Each character in stones is a type of stone you have.
//You want to know how many of the stones you have are also jewels.
//Letters are case sensitive, so "a" is considered a different type of stone from "A".

//func numJewelsInStones(jewels string, stones string) int {
//    var res int
//    uniqs := make(map[int32]int, 0)
//    for _, r := range jewels {
//        uniqs[r]++
//    }
//    for _, r := range stones {
//        if _, exist := uniqs[r]; exist {
//            res++
//        }
//    }
//    return res
//}
