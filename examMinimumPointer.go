package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func distanceToOrigin(p Point) float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func closestPoint(points []Point) Point {
	if len(points) == 0 {
		panic("The points slice is empty")
	}

	closest := points[0]
	minDistance := distanceToOrigin(closest)

	for _, point := range points[1:] { // Start from the second point
		dist := distanceToOrigin(point)
		if dist < minDistance {
			closest = point
			minDistance = dist
		}
	}

	return closest
}

func ExamMinimumPointer() {
	var N int
	fmt.Scan(&N)

	points := []Point{}
	for i := 0; i < N; i++ {
		var x, y int
		fmt.Scan(&x, &y)

		var xxxx Point
		xxxx.x = float64(x)
		xxxx.y = float64(y)
		points = append(points, xxxx)
	}

	closest := closestPoint(points)
	fmt.Println(closest.x, closest.y)
}
