package main

import (
	"fmt"
	"math"
)

func fourth() {
	var students, variants, takenPart, takenSeat int
	fmt.Scan(&students, &variants, &takenPart, &takenSeat)
	first := (takenPart-1)*2 + takenSeat - variants
	second := (takenPart-1)*2 + takenSeat + variants
	takenPartFirst := (first + 1) / 2
	takenPartSecond := (second + 1) / 2
	if first > 0 && (second > students || math.Abs(float64(takenPart)-float64(takenPartFirst)) <
		math.Abs(float64(takenPart)-float64(takenPartSecond))) {
		fmt.Printf("%d %d\n", takenPartFirst, 2-first%2)
	} else if second <= students {
		fmt.Printf("%d %d\n", takenPartSecond, 2-second%2)
	} else {
		fmt.Printf("-1")
	}
}
