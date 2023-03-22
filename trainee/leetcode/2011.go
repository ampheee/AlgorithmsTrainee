package leetcode

//There is a programming language with only four operations and one variable X:
//
//++X and X++ increments the value of the variable X by 1.
//--X and X-- decrements the value of the variable X by 1.
//Initially, the value of X is 0.
//
//Given an array of strings operations containing a list of operations, return the final value of X after performing all the operations.

//func finalValueAfterOperations(operations []string) int {
//    var x = 0
//    for _, operation := range operations {
//        switch operation[0] {
//            case '-': x--
//            case '+': x++
//            case 'X':
//                if operation[1] == '-' {
//                    x--
//                } else {
//                    x++
//                }
//        }
//    }
//    return x
//}
