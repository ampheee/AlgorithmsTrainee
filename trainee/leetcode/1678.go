package leetcode

//You own a Goal Parser that can interpret a string command.
//The command consists of an alphabet of "G", "()" and/or "(al)" in some order.
//The Goal Parser will interpret "G" as the string "G", "()" as the string "o", and "(al)" as the string "al".
//The interpreted strings are then concatenated in the original order.
//Given the string command, return the Goal Parser's interpretation of command.

//func interpret(command string) string {
//    var prev = rune(command[0])
//    var res string
//    for _, r := range command{
//        if prev == '(' && r == ')' {
//            res += "o"
//        } else if r == 'G' {
//            res += "G"
//        } else if prev != '(' && r == ')'{
//            res += "al"
//        }
//        prev = r
//    }
//    return res
//}
