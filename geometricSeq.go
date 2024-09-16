package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//แต่ละตัวจะเท่ากัน
func isArithmeticProgression(arr []int) bool {
	if len(arr) < 2 {
		return false
	}
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

//เลขยกกำลัง ทบไปเรื่อยๆ
func isGeometricProgression(arr []int) bool {
	if len(arr) < 2 {
		return false
	}
	// Avoid division by zero
	if arr[0] == 0 {
		return false
	}
	ratio := arr[1] / arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i-1] == 0 || arr[i]/arr[i-1] != ratio {
			return false
		}
	}
	return true
}

// Function to classify the sequence
func classifySequence(arr []int) string {
	if len(arr) < 2 {
		return "Random"
	}
	if isArithmeticProgression(arr) {
		return "Arithmetic Progression" //เลขคณิต
	} else if isGeometricProgression(arr) {
		return "Geometric Progression" //เรขาคณิต
	} else {
		return "Random"
	}
}

func ExamGeometricSeq() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	arrayOfNumber := scanner.Text()
	_ = arrayOfNumber

	var a = strings.Split(arrayOfNumber, " ")
	values := make([]int, 0, len(a))
	for i := 0; i < len(a); i++ {
		val, _ := strconv.Atoi(a[i])
		values = append(values, val)
	}

	fmt.Println(classifySequence(values))
}
