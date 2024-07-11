package main

import (
	"fmt"
)

// compare อุณหภูมิระหว่าง ฟาเรนไฮต์ กับ เซลเซียส
// problem (input ตัวแรก : ฟาเรนไฮต์ , input ตัวสอง : เซลเซียส)
// input : 5 -10

func ExamCompareTemperature() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var b, t int
		fmt.Scan(&b, &t)

		//สูตร
		tInF := float64(t)*9/5 + 32
		if float64(b) > tInF {
			fmt.Println("Higher")
		} else if float64(b) < tInF {
			fmt.Println("Lower")
		} else {
			fmt.Println("Same")
		}
	}
}
