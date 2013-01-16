package main

import (
	"log"
	"net"
)

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
