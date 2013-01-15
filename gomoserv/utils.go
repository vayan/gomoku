package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"strconv"
)

func Atoi(st string) int {
	ret, _ := strconv.Atoi(st)
	return ret
}

func getClient(ws *websocket.Conn) Connection {
	var connect Connection

	for pl, _ := range players {
		if pl.ws == ws {
			connect = pl
			break
		}
	}
	return (connect)
}

func AffBoard(board [][]int, size int) {
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if board[x][y] == BLACK {
				fmt.Print("B")
			} else if board[x][y] == WHITE {
				fmt.Print("W")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Print("\n")
	}
}

func initBoard(size int) [][]int {

	var board = make([][]int, size)
	BPOW = 0
	WPOW = 0
	Turn = BLACK

	for pl, _ := range players {
		pl.ws.Close()
	}
	players = make(map[Connection]int)

	for x := 0; x < size; x++ {
		board[x] = make([]int, size)
		for y := 0; y < size; y++ {
			board[x][y] = NONE
		}
	}
	return (board)
}
func getinvturn() int {
	if Turn == BLACK {
		return WHITE
	}
	return BLACK
}

func getStringTurn() string {
	if Turn == BLACK {
		return "black"
	}
	return "white"
}

func getStringTurnInv() string {
	if Turn == BLACK {
		return "white"
	}
	return "black"
}

func getStringPl(pl int) string {
	if pl == BLACK {
		return "black"
	} else if pl == WHITE {
		return "white"
	}
	return "OBS"
}

func HandleErrorFatal(er error) bool {
	if er != nil {
		log.Fatal(er)
	}
	return false
}
