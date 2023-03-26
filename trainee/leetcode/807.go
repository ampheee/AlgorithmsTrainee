package leetcode

//There is a city composed of n x n blocks, where each block contains a
//single building shaped like a vertical square prism.
//You are given a 0-indexed n x n integer matrix grid where
//grid[r][c] represents the height of the building located in the block at row r and column c.
//A city's skyline is the the outer contour formed by all the building
//when viewing the side of the city from a distance.
//The skyline from each cardinal direction north, east, south, and west may be different.
//We are allowed to increase the height of any number of buildings
//by any amount (the amount can be different per building). 
//The height of a 0-height building can also be increased.
//However, increasing the height of a building should not affect the city's skyline from any cardinal direction.
//Return the maximum total sum that the height of the
//buildings can be increased by without changing the city's skyline from any cardinal direction.

//func maxIncreaseKeepingSkyline(grid [][]int) int {
//    cols, rows := make([]int, len(grid)), make([]int, len(grid[0]))
//    for i := 0; i < len(grid[0]); i++ {
//        for j := 0; j < len(grid); j++ {
//            rows[i] = int(math.Max(float64(grid[i][j]), float64(rows[i])))
//            cols[i] = int(math.Max(float64(grid[j][i]), float64(cols[i])))
//        }
//    }
//    var res int
//    for i := 0; i < len(grid); i++ {
//        for j := 0; j < len(grid[0]); j++ {
//            res += int(math.Min(float64(cols[j]), float64(rows[i]))) - grid[i][j]
//        }
//    }
//    return res
//}
