package gomoku

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
}

func ws_send(buf string, ws *websocket.Conn) {
	err := websocket.Message.Send(ws, buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("send:%s\n", buf)
}

func ws_recv(ws *websocket.Conn) string {
	var buf string
	//var connect Connection

	err := websocket.Message.Receive(ws, &buf)
	if err != nil {
		// for pl, _ := range players {
		// 	if pl.ws == ws {
		// 		connect = pl
		// 		break
		// 	}
		// }
		fmt.Println(err)
		//delete(players, connect)
	}
	fmt.Printf("recv coord:%s\n", buf)
	return buf
}

func SendRecvCoord(ws *websocket.Conn) {
	slotleft := BLACK
	for pl, _ := range players {
		if pl.player_color == BLACK {
			slotleft = WHITE
		} else {
			slotleft = NONE
		}
	}
	sock_cli := Connection{ws, slotleft, false}
	fmt.Printf("\nNouveau joueurs de type %d\n", slotleft)
	players[sock_cli] = 0

	for {
		var buf string

		buf = ws_recv(ws)

		//check avec le referee
		if buf == "reset" {
			Board = initBoard(GOBANSIZE)

		} else if buf == "getturn" {
			ws_send("turn,"+getStringTurn(), ws)
		} else if buf == "getme" {
			ws_send("me, You are "+getStringPl(getClient(ws).player_color), ws)
		} else if buf == "getscore" {
			ws_send("score, Black : "+strconv.Itoa(BPOW)+" | White : "+strconv.Itoa(WPOW), ws)
		} else {
			mov, win, who := referee(strings.Split(buf, ","), ws)
			if win {
				buf = "win," + getStringPl(who)
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
