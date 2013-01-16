package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"strconv"
)

func GenBoard(size int, stone bool) string {
	board := ""
	color := ""

	for x := 0; x < size; x++ {
		board += "<tr>"
		for y := 0; y < size; y++ {
			if Board[x][y] == BLACK {
				color = "bgblack"
			} else if Board[x][y] == WHITE {
				color = "bgwhite"
			} else {
				color = ""
			}
			if stone == true {
				board += " <td><div class='" + color + " stone pos" + strconv.Itoa(x) + "y" + strconv.Itoa(y) + "'>" + strconv.Itoa(x) + " " + strconv.Itoa(y) + "</div></td> "
			} else {
				board += " <td></td> "
			}
		}
		board += "</tr>"
	}
	return board
}

func getFreeSlot() int {
	slotleft, black, white := NONE, 0, 0

	for pl, _ := range players {
		if pl.player_color == BLACK {
			black = 1
		}
		if pl.player_color == WHITE {
			white = 1
		}
	}

	if black == 0 {
		slotleft = BLACK
	} else if white == 0 && black == 1 {
		slotleft = WHITE
	} else if black == 1 && white == 1 {
		slotleft = NONE
	}
	return slotleft
}

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
