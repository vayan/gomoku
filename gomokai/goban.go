package main

import (
	"fmt"
)

var (
	Board [][]int = initBoard(20)
	Pris          = 0
	Turn          = BLACK
)

const (
	BLACK           = 1
	WHITE           = 2
	NONE            = 0
	NB_ALIGN_TO_WIN = 5
	STONE_TO_Win    = 10
)

func duplicate_board() [][]int {
	var board = make([][]int, 20)

	for x := 0; x < 20; x++ {
		board[x] = make([]int, 20)
		for y := 0; y < 20; y++ {
			board[x][y] = Board[x][y]
		}
	}
	return (board)
}

func change_turn() {
	if Turn == BLACK {
		Turn = WHITE
		return
	}
	Turn = BLACK
}

func rmStone(x int, y int) {
	Board[x][y] = NONE
}

func addStone(x int, y int, color int) {
	Board[x][y] = color
}

func initBoard(size int) [][]int {

	var board = make([][]int, size)

	for x := 0; x < size; x++ {
		board[x] = make([]int, size)
		for y := 0; y < size; y++ {
			board[x][y] = NONE
		}
	}
	return (board)
}

func AffBoard(size int) {
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if Board[x][y] == BLACK {
				fmt.Print("B")
			} else if Board[x][y] == WHITE {
				fmt.Print("W")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Print("\n")
	}
}
