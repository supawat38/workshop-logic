package main

import (
	"bufio"
	"fmt"
	"os"
)

// วงเล็บต้องครบตามที่กำหนด
// problem
// input : ((((())))()()())

func ExamBracketCorrect() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	S := scanner.Text()
	_ = S // to avoid unused error

	if S[:1] == ")" || S[len(S)-1:len(S)] == "(" {
		fmt.Println("false")
		return
	}

	fmt.Println("true") // Write answer to stdout
}
