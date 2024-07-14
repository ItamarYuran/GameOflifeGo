package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	live  = "+ "
	dead  = "  "
	frame = "@ "
)

func main() {

	startGame()
}

func initiate(length, height, percent int) [][]string {
	board := make([][]string, length+2)
	for i := range board {
		board[i] = make([]string, height+2)
	}
	rand.Seed(time.Now().UnixNano())

	for i, row := range board {
		for j := range row {
			if i == 0 || i == length+1 {
				board[i][j] = frame
				continue
			}
			if j == 0 || j == height+1 {
				board[i][j] = frame
				continue
			}
			chance := rand.Intn(100)
			if chance <= percent {
				board[i][j] = dead
			} else {
				board[i][j] = live
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

func getliveneighbors(board [][]string, i int, j int) (count int) {
	for q := -1; q < 2; q++ {
		for p := -1; p < 2; p++ {
			if q == 0 && p == 0 {
				continue
			}
			if board[i+q][j+p] == live {
				count++
			}
		}
	}
	return count
}

func nextboard(board [][]string) (nextboard [][]string) {
	nextboard = make([][]string, len(board))
	for i := range board {
		nextboard[i] = make([]string, len(board[i]))
	}
	for i, row := range board {
		for j := range row {
			if i == 0 || i == len(board)-1 {
				nextboard[i][j] = board[i][j]

				continue
			}
			if j == 0 || j == len(board[0])-1 {
				nextboard[i][j] = board[i][j]
				continue
			}

			neighbors := getliveneighbors(board, i, j)

			if board[i][j] == dead {
				if neighbors == 3 {
					nextboard[i][j] = live
				} else {
					nextboard[i][j] = dead
				}
			}

			if board[i][j] == live {
				if neighbors < 2 || neighbors > 3 {
					nextboard[i][j] = dead
				} else {
					nextboard[i][j] = live
				}
			}

		}
	}
	return nextboard

}

func startGame() {
	var row, col, percent int
	fmt.Printf("What is your preferred length?\n")
	_, err := fmt.Scanln(&row)
	if err != nil {
		fmt.Println("Input isn't valid capara")
	}
	fmt.Printf("What is your preferred height?\n")

	_, err = fmt.Scanln(&col)
	if err != nil {
		fmt.Println("Input isn't valid capara")
	}

	fmt.Printf("What percent of the board should be alive?\n")
	_, err = fmt.Scanln(&percent)
	if err != nil {
		fmt.Println("Input isn't valid capara")
	}
	runGame(row, col, percent)

}

func runGame(length, height, percent int) {

	board := initiate(length, height, percent)

	for {
		clearScreen()
		printBoard(board)
		board = nextboard(board)
		time.Sleep(80 * time.Millisecond)

	}

}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
