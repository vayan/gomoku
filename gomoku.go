package main

import (
	"code.google.com/p/go.net/websocket"
	"html/template"
	"strconv"
)

var Board [][]int = initBoard(GOBANSIZE)

var Turn = BLACK
var BPOW = 0
var WPOW = 0

const (
	BLACK           = 1
	WHITE           = 2
	IA              = 3
	NONE            = 0
	GOBANSIZE       = 20
	NB_ALIGN_TO_WIN = 5
	STONE_TO_Win    = 10
)

type Page struct {
	Title      string
	Board      template.HTML
	BoardClick template.HTML
}

func capture_win() {

}

func send_capture(pow [][]int, ws *websocket.Conn) {
	var buff string
	flag := 0

	if pow != nil {
		for key := range pow {
			if flag == 0 {
				buff = "REM "
			}
			if pow[key] != nil {
				buff += strconv.Itoa(pow[key][0]) + " " + strconv.Itoa(pow[key][1]) + " "
				if Board[pow[key][0]][pow[key][1]] == BLACK {
					WPOW += 1
				} else {
					BPOW += 1
				}
				Board[pow[key][0]][pow[key][1]] = NONE
				if flag == 1 {
					flag = 0
					for pl, _ := range players {
						ws_send(buff, pl.ws)
					}
					continue
				}
				flag++
			}

		}
	}
}
