package main

import (
	"fmt"
	"strings"
)

// split คำตัวหน้า(จากหน้าสุด) และ ตัวหน้า(จากหลังสุด)
// problem
// input : absolutes cars radio

func ExamSplitTextFirstAndLast() {
	var s = "absolutes cars radio"

	var answer = ""
	var lastText = ""
	res1 := strings.Split(s, " ")
	for i := 0; i < len(res1); i++ {
		first := res1[i][:1]
		answer += first

		last := res1[i][len(res1[i])-1:]
		lastText += last
	}

	fmt.Println(answer + Reverse(lastText))
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
