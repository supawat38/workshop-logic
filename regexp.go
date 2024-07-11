package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// เอาเฉพาะตัวอักษร และตัวเลข และเรียงจากน้อยไปมาก (a-Z) (0-9)
// problem
// input : fewgret%sdg234234

func ExamRegexp() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	s := scanner.Text()
	_ = s

	var newString = ""
	var new = []rune(s)
	for i := 0; i < len(new); i++ {
		s := fmt.Sprintf("%c", new[i])
		reg := regexp.MustCompile(`[^a-zA-Z]`)
		if !reg.MatchString(s) {
			newString += s
		}
	}

	var newString2 = ""
	var new2 = []rune(s)
	for i := 0; i < len(new2); i++ {
		s := fmt.Sprintf("%c", new[i])
		reg := regexp.MustCompile(`[^0-9]`)
		if !reg.MatchString(s) {
			newString2 += s
		}
	}

	w2 := SortString(newString) + SortString(newString2)
	fmt.Println("answer :  ", w2)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
