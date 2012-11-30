package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"gomoku"
	"net/http"
)

func main() {
	fmt.Print("\n\n========GEN BOARD========\n")
	gomoku.AffBoard(gomoku.Board, gomoku.GOBANSIZE)

	fmt.Print("\n\n========Start gomoku web server========\n")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/ws", websocket.Handler(gomoku.SendRecvCoord))
	http.HandleFunc("/", gomoku.IndexHandler)

	http.ListenAndServe(":1112", nil)
	fmt.Print("\n\n========Started and listen========\n")
}
