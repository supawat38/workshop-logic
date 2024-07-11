package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// หาสองคำจากข้อความทั้งหมด ที่ถูก sort แล้วมากที่สุด(a) และตัวที่สองคือน้อยที่สุด (z)
// problem
// input : abcd cdab

func ExamMaxMinInString() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	s := scanner.Text()
	_ = s // to avoid unused error

	var textAppend = []string{}
	var newX2 = []rune(s)
	for i := 0; i < len(newX2); i++ {
		textX2 := fmt.Sprintf("%c", newX2[i])
		if i == 0 {
			textAppend = append(textAppend, textX2)
		} else if i == len(newX2)-1 {
			textAppend = append(textAppend, textX2)
		}
	}

	sort.Strings(textAppend)

	if len(textAppend) == 1 {
		fmt.Println(textAppend[0])
		fmt.Println(textAppend[0])
		return
	}

	fmt.Println(textAppend[0])
	fmt.Println(textAppend[1])
}
