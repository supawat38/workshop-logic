package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// พิมพ์คำตั้งต้น และ คำถัดไปเพื่อเช็คว่า สลับคำ กลับคำแล้วต้องได้คำเดิม
// problem
// input : abcd cdab

func ExamRecursiveString() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	s1 := scanner.Text()
	_ = s1 // to avoid unused error

	scanner.Scan()
	s2 := scanner.Text()
	_ = s2 // to avoid unused error

	// s1 = "abcd"
	// s2 = "cdab"

	X1 := strings.ToLower(s1)
	X2 := strings.ToLower(s2)

	var textAppendX2 = []string{}
	var newX2 = []rune(X2)
	for i := 0; i < len(newX2); i++ {
		textX2 := fmt.Sprintf("%c", newX2[i])
		textAppendX2 = append(textAppendX2, textX2)
	}

	var textAnswer bool
	var textCompare string
	for j := 0; j < len(textAppendX2); j++ {
		// fmt.Println(textAppendX2[len(textAppendX2)-1]) //4
		// fmt.Println(textAppendX2[0]) //1
		// fmt.Println(textAppendX2[2]) //2
		// fmt.Println(textAppendX2[3]) //2

		// for w := 0; w < 2; w++{
		// 	textcompare = textAppendX2[len(textAppendX2)-1] +
		// 	textAppendX2[0] +
		// 	textAppendX2[1] +
		// 	textAppendX2[2]

		// 	fmt.Println(textcompare)
		// 	if textcompare == X1 {
		// 		textAnswer = true
		// 	}
		// }

		textCompare = textAppendX2[len(textAppendX2)-1] +
			textAppendX2[0] +
			textAppendX2[1] +
			textAppendX2[2]

		fmt.Println(textCompare)

		if textCompare == X1 {
			textAnswer = true
		}
	}

	fmt.Println(textAnswer)
}
