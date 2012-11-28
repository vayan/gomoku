package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var Board [][]int = initBoard(19)
var Turn = BLACK

const (
	BLACK = 1
	WHITE = 2
	NONE  = 0
)

type Page struct {
	Title      string
	Board      template.HTML
	BoardClick template.HTML
}

func affBoard(board [][]int, size int) {
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if board[x][y] == BLACK {
				fmt.Print("B")
			} else if board[x][y] == WHITE {
				fmt.Print("W")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Print("\n")
	}
}

func initBoard(size int) [][]int {

	var board = make([][]int, size)

	Turn = BLACK

	for x := 0; x < size; x++ {
		board[x] = make([]int, size)
		for y := 0; y < size; y++ {
			board[x][y] = NONE
		}
	}
	return (board)
}

func referee(coord []string) bool {
	x, _ := strconv.Atoi(coord[0])
	y, _ := strconv.Atoi(coord[1])

	if Board[x][y] == NONE {
		Board[x][y] = Turn
		if Turn == BLACK {
			Turn = WHITE
		} else {
			Turn = BLACK
		}
		return (true)
	}
	return false
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

	err := websocket.Message.Receive(ws, &buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("recv coord:%s\n", buf)
	return buf
}

func getStringTurn() string {
	if Turn == BLACK {
		return "white"
	}
	return "black"
}

func sendRecvCoord(ws *websocket.Conn) {
	for {
		var buf string

		buf = ws_recv(ws)

		//check avec le referee
		if buf == "reset" {
			Board = initBoard(19)
		} else {
			if referee(strings.Split(buf, ",")) == true {
				buf += "," + getStringTurn()
			} else {
				buf = "error"
			}
			ws_send(buf, ws)
		}
	}
}

func GenBoard(size int, stone bool) string {
	board := ""
	color := ""

	for x := 0; x < size; x++ {
		board += "<tr>"
		for y := 0; y < size; y++ {
			if Board[x][y] == BLACK {
				color = "bgblack"
			} else if Board[x][y] == WHITE {
				color = "bgwhite"
			} else {
				color = ""
			}
			if stone == true {
				board += " <td><div class='" + color + " stone pos" + strconv.Itoa(x) + "y" + strconv.Itoa(y) + "'>" + strconv.Itoa(x) + "," + strconv.Itoa(y) + "</div></td> "
			} else {
				board += " <td></td> "
			}
		}
		board += "</tr>"
	}
	return board
}

func HandleErrorFatal(er error) bool {
	if er != nil {
		log.Fatal(er)
	}
	return false
}

func loadPage() *Page {
	title := "test"
	board := template.HTML(GenBoard(19, false))
	boardClick := template.HTML(GenBoard(19, true))
	return &Page{Title: title, Board: board, BoardClick: boardClick}
}

func RenderHtml(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	p := loadPage()
	RenderHtml(w, "index", p)
}

func main() {
	fmt.Print("\n\n========GEN BOARD========\n")
	affBoard(Board, 19)

	fmt.Print("\n\n========Start gomoku web server========\n")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/ws", websocket.Handler(sendRecvCoord))
	http.HandleFunc("/", IndexHandler)

	http.ListenAndServe(":1112", nil)
	fmt.Print("\n\n========Started and listen========\n")
}
