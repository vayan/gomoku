package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"strconv"
	"strings"
)

var players = make(map[Connection]int)

type Connection struct {
	ws           *websocket.Conn
	player_color int
	is_ia        bool
	clientip     string
}

func ws_send(buf string, ws *websocket.Conn) {
	err := websocket.Message.Send(ws, buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("send:%s\n", buf)
}

func ws_recv(ws *websocket.Conn) (string, int) {
	var buf string
	erri := 0

	err := websocket.Message.Receive(ws, &buf)
	if err != nil {
		erri = 1
		for pl, _ := range players {
			if pl.ws == ws {
				fmt.Printf("\n*************Deconnexion de %s\n", getStringPl(pl.player_color))
				//pl.ws.Close()
				delete(players, pl)
				break
			}
		}
		fmt.Println(err)
	}
	fmt.Printf("recv :%s\n", buf)
	return buf, erri
}

func sendboard(ws *websocket.Conn) {
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			buf := strconv.Itoa(x) + "," + strconv.Itoa(y)
			if Board[x][y] == BLACK {
				buf += ",black"
				ws_send(buf, ws)
			} else if Board[x][y] == WHITE {
				buf += ",white"
				ws_send(buf, ws)
			}
		}
		fmt.Print("\n")
	}
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

func SendRecvCoord(ws *websocket.Conn) {

	sock_cli := Connection{ws, getFreeSlot(), false, ws.Request().RemoteAddr}
	fmt.Printf("\nNouveau joueurs de type %d\n", sock_cli.player_color)
	sendboard(ws)
	players[sock_cli] = 0

	for {
		var buf string
		var erri int

		if buf, erri = ws_recv(ws); erri == 1 {
			return
		}

		//check avec le referee
		if buf == "reset" {
			Board = initBoard(GOBANSIZE)
			return

		} else if buf == "getturn" {
			ws_send("turn,"+getStringTurn(), ws)
		} else if buf == "getme" {
			ws_send("me, You are "+getStringPl(getClient(ws).player_color), ws)
		} else if buf == "getscore" {
			ws_send("score, Black : "+strconv.Itoa(BPOW)+" | White : "+strconv.Itoa(WPOW), ws)
		} else {
			coord := strings.Split(buf, ",")
			if len(coord) > 1 {
				mov, win, who := referee(coord, ws)
				if win {
					buf = "win," + getStringPl(who)
					for pl, _ := range players {
						ws_send(buf, pl.ws)
					}
				} else if mov == true {
					buf += "," + getStringTurnInv()
					for pl, _ := range players {
						ws_send(buf, pl.ws)
					}
				} else {
					buf = "error"
				}
				ws_send(buf, ws)
				AffBoard(Board, GOBANSIZE)
			}
		}

	}
}
