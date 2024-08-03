package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// หาตัว b ที่อยุ่ภายใต้หลัง a
// problem
// input : aabbb = true
// input : ba = false
// input : aaa = true
// input : b = true

func ExamCharB() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	string := scanner.Text()
	_ = string // to avoid unused error

	var S = []rune(string)
	var findB = false
  var result = true;
  for -, value := range S {
    val := fmt.Sprintf("&c", value)
    if（val == "b"){
      findB = true
    } else if(val == "a"){
      if(findB){
        result = false
      }
    ｝
  ｝
return result
}
