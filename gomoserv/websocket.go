package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
)

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

func HandleWebSocket(ws *websocket.Conn) {

	sock_cli := Connection{ws, nil, getFreeSlot(), false, ws.Request().RemoteAddr}
	log.Printf("\nNouveau joueurs de type %d\n", sock_cli.player_color)
	sendboard(sock_cli)
	send(RULES_ST, sock_cli)
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
