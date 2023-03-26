package leetcode

//You are given an array points where points[i] = [xi, yi] is the coordinates of the ith point on a 2D plane.
//Multiple points can have the same coordinates.
//You are also given an array queries where queries[j] = [xj, yj, rj]
//describes a circle centered at (xj, yj) with a radius of rj.
//For each query queries[j], compute the number of points inside the jth circle.
//Points on the border of the circle are considered inside.

//Return an array answer, where answer[j] is the answer to the jth query.

//func countPoints(points [][]int, queries [][]int) []int {
//    var res = make([]int, 0, 0)
//    for _, circle := range queries {
//        var count = 0
//        for _, point := range points {
//            if pow(point[0] - circle[0]) + pow(point[1] - circle[1]) <= pow(circle[2]) {
//                    count++
//                }
//        }
//        res = append(res, count)
//    }
//    return res
//}
//
//func pow(a int) int {
//    return a * a
//}
