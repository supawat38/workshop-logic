package main

import (
	"fmt"
)

func ExamSlice() {
	// อ่านค่าข้อมูลเข้า
	var N, M int
	fmt.Scan(&N)

	// อ่าน N จำนวนเต็มเข้าไปในอาร์เรย์
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Scan(&M)

	// วนลูปผ่านทุกช่วง
	for i := 0; i < M; i++ {
		var startRange, endRange int
		fmt.Scan(&startRange, &endRange)

		// คำนวณผลรวมของค่าในช่วงที่กำหนด
		sum := 0
		for j := startRange; j <= endRange; j++ {
			sum += arr[j]
		}

		// แสดงผลรวมสำหรับช่วงปัจจุบัน
		fmt.Println(sum)
	}
}
