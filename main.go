package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net"
	"net/http"
)

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
