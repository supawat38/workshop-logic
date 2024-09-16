package main

import "fmt"

func ProcessBalanceBank(recipients string, amounts []int) []int {
	minA := 0
	minB := 0

	balA := 0
	balB := 0

	for i := 0; i < len(recipients); i++ {
		if recipients[i] == 'A' {
			balA += amounts[i]
			balB -= amounts[i]
			if balB < minB {
				minB = balB
			}
			fmt.Println("A balA (+): ", balA)
			fmt.Println("A balB (-): ", balB)
		} else if recipients[i] == 'B' {
			balB += amounts[i]
			balA -= amounts[i]
			if balA < minA {
				minA = balA
			}
			fmt.Println("B balA (-): ", balA)
			fmt.Println("B balB (+): ", balB)
		}
	}

	return []int{-minA, -minB}
}

func ExamBalanceBank() {
	var result = ProcessBalanceBank("BAABA", []int{2, 4, 1, 1, 2})
	fmt.Println(result)
}

//----------------------
// 			   A     B
// initial     2     4
// A 2 B       0     6
// B 4 A       4     2
// B 1 A       5     1
// A 1 B       4     2
// B 2 A       6     0

//----------------------
//             A     B
// initial     0     0
// A 2 B      -2     2
// B 4 A       2    -2
// B 1 A       3    -3
// A 1 B       2    -2
// B 2 A       4    -4

// answer : [2,4]

