package main

import (
	"code.google.com/p/go.net/websocket"
	"html/template"
	"log"
	"net"
	"strings"
)

var (
	Board      [][]int = initBoard(GOBANSIZE)
	Turn               = BLACK
	BPOW               = 0
	WPOW               = 0
	Mode               = UNKNOWN
	DOUBLE_3           = 0
	BREAKING_5         = 0
	TIMEOUT            = 0
	players            = make(map[Connection]int)
	tm                 = 0
	RULES_ST           = "RULES 0 0 0"
)

const (
	UNKNOWN         = -1
	PVP             = 1
	PVE             = 2
	BLACK           = 1
	WHITE           = 2
	IA              = 3
	NONE            = 0
	GOBANSIZE       = 20
	NB_ALIGN_TO_WIN = 5
	STONE_TO_Win    = 10
	W_CAPTURE       = 1
	W_FIVEALIGN     = 2
	W_RULEERR       = 3
	W_TIMEOUT       = 4
)

type Page struct {
	Title      string
	Board      template.HTML
	BoardClick template.HTML
}

type Connection struct {
	ws           *websocket.Conn
	s            net.Conn
	player_color int
	is_ia        bool
	clientip     string
}

func engine(msg_cl string, con Connection) int {
	buff := strings.Split(msg_cl, " ")

	buf := buff[0]

	if con.player_color == NONE {
		return 1
	}

	//check le mode
	if Mode == UNKNOWN {
		if buf == "MODE" {
			if buff[1] == "pve" {
				Mode = PVE
				con.player_color = BLACK
			}
			if buff[1] == "pvp" {
				Mode = PVP
			}
		}
		return 1
	}

	//check avec le referee
	switch buf {
	case "reset":
		for pl, _ := range players {
			delete(players, pl)
		}
		players = make(map[Connection]int)
		Board = initBoard(GOBANSIZE)
		DOUBLE_3 = 0
		BREAKING_5 = 0
		TIMEOUT = 0
		return -1

	case "RULES":
		DOUBLE_3 = Atoi(buff[1])
		BREAKING_5 = Atoi(buff[2])
		TIMEOUT = Atoi(buff[3])
		RULES_ST = msg_cl
		for pl, _ := range players {
			send(msg_cl, pl)
		}
	case "GETCOLOR":
		send("COLOR "+getStringPl(con.player_color), con)
	case "GETTURN":
		send("TURN "+getStringTurn(), con)
	case "CONNECT":
		if buff[1] == "CLIENT" {
			log.Print("IA Connected")
			con.is_ia = true
			con.player_color = WHITE
			send(RULES_ST, con)
			if Turn == WHITE {
				send("YOURTURN", con)
			}
		}
	case "PLAY":
		tm++
		coord := []string{buff[1], buff[2]}
		if len(coord) > 1 {
			mov, win, _ := referee(coord, con)
			if win {
				send_win(con, W_FIVEALIGN)
				return -1
			} else if mov == true {
				buf = "ADD " + buff[1] + " " + buff[2]
				for pl, _ := range players {
					send(buf, pl)
					if pl.player_color == Turn {
						send("YOURTURN", pl)
						if Mode == PVE && Turn == WHITE && TIMEOUT > 0 {
							go timeout_to_death()
						}
					}
				}
			} else {
				if Mode == PVE && con.player_color == WHITE {
					for pl, _ := range players {
						if pl.player_color == BLACK {
							send_win(pl, W_RULEERR)
							return -1
						}
					}
				}
			}
			//send(buf, ws)
			AffBoard(Board, GOBANSIZE)
		}

	}
	return 1
}
