package main

import "fmt"

func ExamChannel() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{10, 20, 30}

	sumArr1 := make(chan int)
	sumArr2 := make(chan int)

	go func(arr1 []int, sum chan int) {
		var result int
		for i := range arr1 {
			result += arr1[i]
		}
		sum <- result
		// Don't forget to close a channel after assigned
		close(sum)
	}(arr1, sumArr1)

	go func(arr2 []int, sum chan int) {
		var result int
		for i := range arr2 {
			result += arr2[i]
		}
		sum <- result
		// Don't forget to close a channel after assigned
		close(sum)
	}(arr2, sumArr2)

	sum1 := <-sumArr1
	sum2 := <-sumArr2

	fmt.Println("sum1 : ", sum1)
	fmt.Println("sum2 : ", sum2)
}
