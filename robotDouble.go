package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// พิมพ์ข้อความทิศทาง และ จะออกมาเป็นเครื่องหมาย ถ้าทิศทางเดิม จะมีตัวเลขต่อท้าย
// problem
// input : up up down left right

func keyword(order string) string {
	switch order {
	case "up":
		return "^"
	case "down":
		return "v"
	case "left":
		return "<"
	case "right":
		return ">"
	default:
		return ""
	}
}

func compressOrders(orders []string) string {
	var result strings.Builder
	currentOrder := keyword(orders[0])
	count := 1

	for i := 1; i < len(orders); i++ {
		orderSymbol := keyword(orders[i])

		//if same -> count +
		if orderSymbol == currentOrder {
			count++
		} else {
			if count > 1 {
				result.WriteString(fmt.Sprintf("%s%d", currentOrder, count))
			} else {
				result.WriteString(currentOrder)
			}
			currentOrder = orderSymbol
			count = 1
		}
	}

	if count > 1 {
		result.WriteString(fmt.Sprintf("%s%d", currentOrder, count))
	} else {
		result.WriteString(currentOrder)
	}

	return result.String()
}

func ExamRobotDoubleText() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	orders := scanner.Text()
	_ = orders

	InputArray := strings.Split(orders, " ")
	compressedOrders := compressOrders(InputArray)
	fmt.Println(compressedOrders)
}
