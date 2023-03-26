package leetcode

// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
//
// You have to rotate the image in-place, which means you have to modify
// the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
//func rotate(matrix [][]int) {
//	for i := 0; i < len(matrix); i++ {
//		for j := i; j < len(matrix); j++ {
//			swap(&matrix[i][j], &matrix[j][i])
//		}
//	}
//	for i := range matrix {
//		reverse(matrix[i])
//	}
//	fmt.Println(matrix)
//}
//
//func reverse(matrixString []int) {
//	n := len(matrixString)
//	for i := 0; i < n/2; i++ {
//		j := n - i - 1
//		swap(&matrixString[i], &matrixString[j])
//	}
//}
//
//func swap(a, b *int) {
//	temp := *a
//	*a = *b
//	*b = temp
//}
