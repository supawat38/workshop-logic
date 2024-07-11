package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// สลับคำจากหลังมาหน้า ไม่เอาตัวเลข และแปลงเป็นตัวใหญ่
// problem
// input : 123213dsfdsardfsv

func ExamReverse() {
	//reverse
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	cypher := scanner.Text()
	_ = cypher // to avoid unused error

	var newString = ""
	var new = []rune(cypher)
	for i := 0; i < len(new); i++ {
		s := fmt.Sprintf("%c", new[i])
		match, _ := regexp.MatchString("[0-9]", s)
		if !match {
			newString += strings.ToUpper(s)
		}
	}

	var result = ""
	for _, v := range newString {
		result = string(v) + result
	}

	fmt.Println("answer : ", result)
}
