package main

import "fmt"

func eight() {
	var count int
	var x, y, xMin, yMin, xMax, yMax float64
	fmt.Scan(&count)
	if count == 1 {
		fmt.Scan(&x, &y)
		fmt.Printf("%.0f %.0f %.0f %.0f\n", x, y, x, y)
		return
	}
	for i := 0; i < count; i++ {
		fmt.Scan(&x, &y)
		if i == 0 {
			yMin = y
			xMin = x
			xMax = x
			yMax = y
		}
		if y <= yMin {
			yMin = y
		}
		if x <= xMin {
			xMin = x
		}
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}
	fmt.Printf("%.0f %.0f %.0f %.0f\n", xMin, yMin, xMax, yMax)
}
