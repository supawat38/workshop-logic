package main

import (
	"fmt"
)

// หาทิศทางที่เดินซ้ำบ่อยสุด
// problem
// input : ^><>^<v<><>

type Position struct {
	x, y int
}

func robotMaxVisits(commands string) int {
	// Initialize variables
	// directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} // N, S, W, E
	currentPos := Position{0, 0}
	visited := make(map[Position]int)
	maxVisits := 0

	// Update the visited count for the starting position
	visited[currentPos]++

	// Simulate robot movement based on commands
	//N = เหนือ S = ใต้ W = ตะวันตก E = ตะวันออก
	for _, command := range commands {
		switch command {
		case '^':
			currentPos.y++
		case '>':
			currentPos.x++
		case 'v':
			currentPos.y--
		case '<':
			currentPos.x--
		}

		// Update visited count for the current position
		visited[currentPos]++

		// Update maxVisits if needed
		if visited[currentPos] > maxVisits {
			maxVisits = visited[currentPos]
		}
	}

	return maxVisits
}

func ExamRobot() {
	commands := "^><>^<v<><>"
	maxVisits := robotMaxVisits(commands)
	fmt.Println("Largest number of visits to the same square:", maxVisits)
}
