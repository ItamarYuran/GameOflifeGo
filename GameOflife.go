package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	board := initiate(10, 10, 80)
	printBoard(board)
}

func initiate(length, height, percent int) [][]string {
	board := make([][]string, length+2)
	for i := range board {
		board[i] = make([]string, height+2)
	}
	rand.Seed(time.Now().UnixNano())

	for i, row := range board {
		for j, _ := range row {
			if i == 0 || i == length+1 {
				board[i][j] = "# "
				continue
			}
			if j == 0 || j == height+1 {
				board[i][j] = "# "
				continue
			}
			chance := rand.Intn(100)
			if chance <= percent {
				board[i][j] = "  "
			} else {
				board[i][j] = "* "
			}

		}
	}
	return board

}

func printBoard(board [][]string) {
	for _, row := range board {
		for _, val := range row {
			fmt.Printf("%v", val)
		}
		fmt.Println()

	}
}
