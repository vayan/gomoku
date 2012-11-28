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

var Board [][]int = initBoard(GOBANSIZE)
var Turn = BLACK
var BPOW = 0
var WPOW = 0

const (
	BLACK           = 1
	WHITE           = 2
	NONE            = 0
	GOBANSIZE       = 20
	NB_ALIGN_TO_WIN = 5
	STONE_TO_Win    = 10
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

func capture_win() {

}

func send_capture(pow [][]int, ws *websocket.Conn) {
	var buff string

	if pow != nil {
		for key := range pow {
			if pow[key] != nil {
				buff = strconv.Itoa(pow[key][0]) + "," + strconv.Itoa(pow[key][1]) + ",pow"
				ws_send(buff, ws)
				if Board[pow[key][0]][pow[key][1]] == BLACK {
					WPOW += 1
				} else {
					BPOW += 1
				}
				Board[pow[key][0]][pow[key][1]] = NONE
			}

		}
	}
}

func capture(coord []int) [][]int {
	x := coord[0]
	y := coord[1]
	i := 0

	var capt = make([][]int, 19)
	to_capt := BLACK
	if Board[x][y] == BLACK {
		to_capt = WHITE
	}

	// hori
	if x-3 >= 0 && Board[x-1][y] == to_capt && Board[x-2][y] == to_capt && Board[x-3][y] == Board[x][y] {
		capt[i] = []int{x - 1, y}
		capt[i+1] = []int{x - 2, y}
		i += 2
	}
	if x+3 <= 19 && Board[x+1][y] == to_capt && Board[x+2][y] == to_capt && Board[x+3][y] == Board[x][y] {
		capt[i] = []int{x + 1, y}
		capt[i+1] = []int{x + 2, y}
		i += 2
	}

	//verti
	if y+3 <= 19 && Board[x][y+1] == to_capt && Board[x][y+2] == to_capt && Board[x][y+3] == Board[x][y] {
		capt[i] = []int{x, y + 1}
		capt[i+1] = []int{x, y + 2}
		i += 2
	}
	if y-3 >= 0 && Board[x][y-1] == to_capt && Board[x][y-2] == to_capt && Board[x][y-3] == Board[x][y] {
		capt[i] = []int{x, y - 1}
		capt[i+1] = []int{x, y - 2}
		i += 2
	}

	// /
	if x-3 >= 0 && y+3 <= 19 && Board[x-1][y+1] == to_capt && Board[x-2][y+2] == to_capt && Board[x-3][y+3] == Board[x][y] {
		capt[i] = []int{x - 1, y + 1}
		capt[i+1] = []int{x - 2, y + 2}
		i += 2
	}
	if x+3 <= 19 && y-3 >= 0 && Board[x+1][y-1] == to_capt && Board[x+2][y-2] == to_capt && Board[x+3][y-3] == Board[x][y] {
		capt[i] = []int{x + 1, y - 1}
		capt[i+1] = []int{x + 2, y - 2}
		i += 2
	}

	// \
	if x-3 >= 0 && y-3 >= 0 && Board[x-1][y-1] == to_capt && Board[x-2][y-2] == to_capt && Board[x-3][y-3] == Board[x][y] {
		capt[i] = []int{x - 1, y - 1}
		capt[i+1] = []int{x - 2, y - 2}
		i += 2
	}
	if x+3 <= 19 && y+3 <= 19 && Board[x+1][y+1] == to_capt && Board[x+2][y+2] == to_capt && Board[x+3][y+3] == Board[x][y] {
		capt[i] = []int{x - 1, y - 1}
		capt[i+1] = []int{x - 2, y - 2}
		i += 2
	}
	if i != 0 {
		return capt
	}
	return nil
}

func check_align(coord []int) bool {
	hor := 0
	vert := 0
	tl := 1
	tr := 1
	x := coord[0]
	y := coord[1]

	//check hor
	for i := y; (i >= 0) && Board[x][i] == Board[x][y]; i-- {
		hor++
	}
	for i := y; (i <= 19) && Board[x][i] == Board[x][y]; i++ {
		hor++
	}
	for i := x; (i <= 19) && Board[i][y] == Board[x][y]; i++ {
		vert++
	}
	for i := x; (i >= 0) && Board[i][y] == Board[x][y]; i-- {
		vert++
	}
	// check \
	for x, y = coord[0]-1, coord[1]-1; x >= 0 && y >= 0 && Board[x][y] == Board[coord[0]][coord[1]]; {
		tl++
		x--
		y--
	}
	for x, y = coord[0]+1, coord[1]+1; x <= 19 && y <= 19 && Board[x][y] == Board[coord[0]][coord[1]]; {
		tl++
		x++
		y++
	}

	//check /
	for x, y = coord[0]-1, coord[1]+1; x >= 0 && y <= 19 && Board[x][y] == Board[coord[0]][coord[1]]; {
		tl++
		x--
		y++
	}
	for x, y = coord[0]+1, coord[1]-1; x <= 19 && y >= 0 && Board[x][y] == Board[coord[0]][coord[1]]; {
		tl++
		x++
		y--
	}

	if hor > NB_ALIGN_TO_WIN || vert > NB_ALIGN_TO_WIN || tl >= NB_ALIGN_TO_WIN || tr >= NB_ALIGN_TO_WIN {
		return true
	}
	return false
}

func check_win(coord []int) bool {
	if check_align(coord) == true {
		return true
	}
	return false
}

func referee(coord []string, ws *websocket.Conn) (bool, bool) {

	x, _ := strconv.Atoi(coord[0])
	y, _ := strconv.Atoi(coord[1])

	coordint := []int{x, y}

	if Board[x][y] == NONE {
		Board[x][y] = Turn
		if Turn == BLACK {
			Turn = WHITE
		} else {
			Turn = BLACK
		}
		send_capture(capture(coordint), ws)
		return true, check_win(coordint)
	}
	return false, false
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
			Board = initBoard(GOBANSIZE)
		} else {
			mov, win := referee(strings.Split(buf, ","), ws)
			if win {
				buf = "win"
			} else if mov == true {
				buf += "," + getStringTurn()
			} else {
				buf = "error"
			}
			ws_send(buf, ws)
			affBoard(Board, GOBANSIZE)
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
	boardClick := template.HTML(GenBoard(GOBANSIZE, true))
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
	affBoard(Board, GOBANSIZE)

	fmt.Print("\n\n========Start gomoku web server========\n")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/ws", websocket.Handler(sendRecvCoord))
	http.HandleFunc("/", IndexHandler)

	http.ListenAndServe(":1112", nil)
	fmt.Print("\n\n========Started and listen========\n")
}
