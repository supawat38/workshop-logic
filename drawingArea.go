package main

import "fmt"

// problem
// width , height , x1 , y1 , x2 , y2
// 30 10 5 5 29 8
// 10 5 3 2 3 2
// 20 8 -10 -8 -1 8
// 30 10 15 5 45 15

func ExamDrawingArea() {
	var width, height int
	fmt.Scan(&width, &height)

	var x1, y1 int
	fmt.Scan(&x1, &y1)

	var x2, y2 int
	fmt.Scan(&x2, &y2)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (j == x1 || j == x2) && i >= y1 && i <= y2 {
				fmt.Print("#")
			} else if (i == y1 || i == y2) && j >= x1 && j <= x2 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
