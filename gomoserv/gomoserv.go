package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net"
	"net/http"
	"time"
)

func timeout_to_death() {
	log.Printf("START TIMEOUT NB %d AI", tm)
	id := tm
	go func() { select {} }()
	select {
	case <-time.After(time.Duration(TIMEOUT) * time.Millisecond):
		if Turn == WHITE && id == tm {
			log.Printf("AI TIMEOUT ON %d (%d)", id, tm)
			for pl, _ := range players {
				if pl.player_color == BLACK {
					send_win(pl, W_TIMEOUT)
					return
				}
			}
			return
		}
	}
	tm++
}

func main() {
	log.Print("\n\n========GEN BOARD========\n")
	AffBoard(Board, GOBANSIZE)

	lis, _ := net.Listen("tcp", ":1113")
	log.Print("== Listen socket 1113 ==")

	go func() {
		for {
			con, err := lis.Accept()
			if err != nil {
				continue
			}
			go HandleSocket(con)
		}
	}()

	go func() {
		log.Print("== Start gomoku web server on 1112 ==")
		http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
		http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
		http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
		http.Handle("/ws", websocket.Handler(HandleWebSocket))
		http.HandleFunc("/", IndexHandler)

		http.ListenAndServe(":1112", nil)
	}()

	select {}
}
