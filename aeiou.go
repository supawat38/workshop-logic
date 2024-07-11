package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// แปลงคำ ถ้าเจอ aeiou ให้เป็น c นอกนั้นเป็น v
// problem
// input : qesadfdsfp

func ExamAeiou() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	string := scanner.Text()
	_ = string // to avoid unused error

	// var string = "Hello World!"
	var new = []rune(string)
	var newText = ""
	var result = ""
	for i := 0; i < len(new); i++ {
		s := fmt.Sprintf("%c", new[i])
		switch s {
		case " ":
			newText = " "
		default:
			match, _ := regexp.MatchString("[A-z]", s)
			if match {
				MatchString, _ := regexp.MatchString("[aeiou]", s)
				if MatchString {
					if strings.ToUpper(s) == s {
						newText = "C"
					} else {
						newText = "c"
					}
				} else {
					if strings.ToUpper(s) == s {
						newText = "V"
					} else {
						newText = "v"
					}
				}
			} else {
				newText = s
			}
		}
		result = result + newText
	}
	fmt.Println(result)
}
