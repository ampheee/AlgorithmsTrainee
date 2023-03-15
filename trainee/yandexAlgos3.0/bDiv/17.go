package main

import "fmt"

func seventeenth() {
	var queue1, queue2 = make([]int, 5), make([]int, 5)
	for j := 0; j < 5; j++ {
		fmt.Scan(&queue1[j])
	}
	for j := 0; j < 5; j++ {
		fmt.Scan(&queue2[j])
	}
	var counter = 0
	for counter < 1000000 {
		var firstCard = queue1[0]
		var secondCard = queue2[0]
		queue1 = queue1[1:]
		queue2 = queue2[1:]
		if firstCard == 0 && secondCard == 9 {
			queue1 = append(queue1, firstCard, secondCard)
		} else if secondCard == 0 && firstCard == 9 {
			queue2 = append(queue2, firstCard, secondCard)
		} else if firstCard > secondCard {
			queue1 = append(queue1, firstCard, secondCard)
		} else if firstCard < secondCard {
			queue2 = append(queue2, firstCard, secondCard)
		}
		counter++
		if len(queue1) == 0 {
			fmt.Println("second", counter)
			return
		}
		if len(queue2) == 0 {
			fmt.Println("first", counter)
			return
		}
	}
	fmt.Println("botva")
}
