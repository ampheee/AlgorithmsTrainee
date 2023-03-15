package main

import (
	"fmt"
	"os"
)

func sixteenth() {
	var queue []int
	for {
		var choice string
		fmt.Scan(&choice)
		switch choice {
		case "push":
			var temp int
			fmt.Scan(&temp)
			queue = append(queue, temp)
			fmt.Println("ok")
		case "front":
			if len(queue) > 0 {
				fmt.Println(queue[0])
			} else {
				fmt.Println("error")
			}
		case "size":
			fmt.Println(len(queue))
		case "pop":
			if len(queue) > 0 {
				var elem = queue[0]
				queue = queue[1:]
				fmt.Println(elem)
			} else {
				fmt.Println("error")
			}
		case "clear":
			for len(queue) != 0 {
				_ = queue[0]
				queue = queue[1:]
			}
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			os.Exit(0)
		}
	}
}
