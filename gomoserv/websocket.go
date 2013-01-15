package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net"
	"strconv"
	"strings"
)

var players = make(map[Connection]int)

type Connection struct {
	ws           *websocket.Conn
	s            net.Conn
	player_color int
	is_ia        bool
	clientip     string
}

func send(buf string, c Connection) {
	log.Printf("send : '%s'", buf)
	if c.ws != nil {
		ws_send(buf, c.ws)
		return
	}
	c.s.Write([]byte(buf + "\n"))
}

func ws_send(buf string, ws *websocket.Conn) {
	err := websocket.Message.Send(ws, buf)
	if err != nil {
		log.Println(err)
	}

}

func ws_recv(ws *websocket.Conn) (string, int) {
	var buf string
	erri := 0

	err := websocket.Message.Receive(ws, &buf)
	if err != nil {
		erri = 1
		for pl, _ := range players {
			if pl.ws == ws {
				log.Printf("\n*************Deconnexion de %s\n", getStringPl(pl.player_color))
				//pl.ws.Close()
				delete(players, pl)
				break
			}
		}
		log.Println(err)
	}
	log.Printf("WS Receive : '%s'", buf)
	return buf, erri
}

func sendboard(s Connection) {

	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			buf := strconv.Itoa(x) + "," + strconv.Itoa(y)
			if Board[x][y] != NONE {
				buf = "ADD " + strconv.Itoa(x) + " " + strconv.Itoa(y)
				send(buf, s)
			}
		}
		log.Print("\n")
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

func engine(msg_cl string, con Connection) int {
	buff := strings.Split(msg_cl, " ")

	buf := buff[0]

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
		return -1

	//case "getturn":
	//  send("turn "+getStringTurn(), ws)
	//case "getme":
	//send("me You are "+getStringPl(getClient(ws).player_color), ws)
	//case "getscore":
	//send("score, Black : "+strconv.Itoa(BPOW)+" | White : "+strconv.Itoa(WPOW), ws)
	case "GETCOLOR":
		send("COLOR "+getStringPl(con.player_color), con)
	case "GETTURN":
		send("TURN "+getStringTurn(), con)
	case "CONNECT CLIENT":
		log.Print("IA Connected")
		con.is_ia = true
		con.player_color = WHITE
		if Turn == WHITE {
			send("YOURTURN", con)
		}
	case "PLAY":
		coord := []string{buff[1], buff[2]}
		if len(coord) > 1 {
			mov, win, _ := referee(coord, con)
			if win {
				send("WIN FIVEALIGN", con)
				for pl, _ := range players {
					if pl != con {
						send("LOSE FIVEALIGN", pl)
					}
				}
			} else if mov == true {
				buf = "ADD " + buff[1] + " " + buff[2]
				for pl, _ := range players {
					send(buf, pl)
					if pl.player_color == Turn {
						send("YOURTURN", pl)
					}
				}
			} else {
				buf = "error"
			}
			//send(buf, ws)
			AffBoard(Board, GOBANSIZE)
		}

	}
	return 1
}

func HandleSocket(con net.Conn) {
	var data = make([]byte, 70)

	log.Printf("=== New Connection received from: %s \n", con.RemoteAddr())
	sock_cli := Connection{nil, con, getFreeSlot(), false, ""}
	log.Printf("\nNouveau joueurs de type %d\n", sock_cli.player_color)
	sendboard(sock_cli)
	players[sock_cli] = 0
	for {
		n, err := con.Read(data)
		if err != nil {
			delete(players, sock_cli)
			log.Println(err)
			return
		}
		buff := string(data[0 : n-1])
		log.Printf("SKT Receive '%s'", buff)
		if engine(buff, sock_cli) == -1 {
			return
		}
		//con.Write(data)
	}
	//log.Println("Data send by client: " + response
}

func HandleWebSocket(ws *websocket.Conn) {

	sock_cli := Connection{ws, nil, getFreeSlot(), false, ws.Request().RemoteAddr}
	log.Printf("\nNouveau joueurs de type %d\n", sock_cli.player_color)
	sendboard(sock_cli)
	players[sock_cli] = 0

	for {
		msg_cl, erri := ws_recv(ws)
		if erri == 1 {
			return
		}
		if engine(msg_cl, sock_cli) == -1 {
			return
		}
	}
}
