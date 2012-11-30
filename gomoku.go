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
	BPOW = 0
	WPOW = 0
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

func breakable_opos(coord []int) bool {
	x := coord[0]
	y := coord[1]
	to_capt := BLACK

	if Board[x][y] == BLACK {
		to_capt = WHITE
	}

	if (x-1 >= 0 && x+1 <= 19 && Board[x-1][y] == to_capt && Board[x+1][y] == Board[x][y]) ||
		(x-1 >= 0 && x+1 <= 19 && Board[x-1][y] == Board[x][y] && Board[x+1][y] == to_capt) ||
		(y+1 <= 19 && y-1 >= 0 && Board[x][y+1] == to_capt && Board[x][y-1] == Board[x][y]) ||
		(y+1 <= 19 && y-1 >= 0 && Board[x][y+1] == Board[x][y] && Board[x][y-1] == to_capt) ||
		(x-1 >= 0 && y+1 <= 19 && x+1 <= 19 && y-1 >= 0 && Board[x-1][y+1] == Board[x][y] && Board[x+1][y-1] == to_capt) ||
		(x-1 >= 0 && y+1 <= 19 && x+1 <= 19 && y-1 >= 0 && Board[x-1][y+1] == to_capt && Board[x+1][y-1] == Board[x][y]) ||
		(x-1 >= 0 && y-1 >= 0 && x+1 <= 19 && y+1 <= 19 && Board[x-1][y-1] == Board[x][y] && Board[x+1][y+1] == to_capt) ||
		(x-1 >= 0 && y-1 >= 0 && x+1 <= 19 && y+1 <= 19 && Board[x-1][y-1] == to_capt && Board[x+1][y+1] == Board[x][y]) {
		return true
	}
	return false
}

func breakable_same(coord []int) bool {
	x := coord[0]
	y := coord[1]
	to_capt := BLACK

	if Board[x][y] == BLACK {
		to_capt = WHITE
	}

	if (x-2 >= 0 && Board[x-1][y] == Board[x][y] && Board[x-2][y] == to_capt) ||
		(x+2 <= 19 && Board[x+1][y] == Board[x][y] && Board[x+2][y] == to_capt) ||
		(y+2 <= 19 && Board[x][y+1] == Board[x][y] && Board[x][y+2] == to_capt) ||
		(y-2 >= 0 && Board[x][y-1] == Board[x][y] && Board[x][y-2] == to_capt) ||
		(x-2 >= 0 && y+2 <= 19 && Board[x-1][y+1] == Board[x][y] && Board[x-2][y+2] == to_capt) ||
		(x+2 <= 19 && y-2 >= 0 && Board[x+1][y-1] == Board[x][y] && Board[x+2][y-2] == to_capt) ||
		(x-2 >= 0 && y-2 >= 0 && Board[x-1][y-1] == Board[x][y] && Board[x-2][y-2] == to_capt) ||
		(x+2 <= 19 && y+2 <= 19 && Board[x+1][y+1] == Board[x][y] && Board[x+2][y+2] == to_capt) {
		return true
	}
	return false
}

func breakable(coord []int) bool {
	if breakable_same(coord) {
		return true
	} else if breakable_opos(coord) {
		return true
	}
	return false
}

func check_all_three_free(coord []int, player int) int {
	x := coord[0]
	y := coord[1]

	nb_free := 0

	for i := y; (i >= 0) && Board[x][i] == Board[x][y]; i-- {
		if create_three_free([]int{x, i}, player) >= 1 {
			nb_free++
			break
		}
	}
	for i := y; (i <= 19) && Board[x][i] == Board[x][y]; i++ {
		if create_three_free([]int{x, i}, player) >= 1 {
			nb_free++
			break
		}

	}
	for i := x; (i <= 19) && Board[i][y] == Board[x][y]; i++ {
		if create_three_free([]int{i, y}, player) >= 1 {
			nb_free++
			break
		}

	}
	for i := x; (i >= 0) && Board[i][y] == Board[x][y]; i-- {
		if create_three_free([]int{i, y}, player) >= 1 {
			nb_free++
			break
		}

	}

	for x, y = coord[0]-1, coord[1]-1; x >= 0 && y >= 0 && Board[x][y] == Board[coord[0]][coord[1]]; {
		if create_three_free([]int{x, y}, player) >= 1 {
			nb_free++
			break
		}

		x--
		y--
	}
	for x, y = coord[0]+1, coord[1]+1; x <= 19 && y <= 19 && Board[x][y] == Board[coord[0]][coord[1]]; {
		if create_three_free([]int{x, y}, player) >= 1 {
			nb_free++
			break
		}

		x++
		y++
	}

	for x, y = coord[0]-1, coord[1]+1; x >= 0 && y <= 19 && Board[x][y] == Board[coord[0]][coord[1]]; {
		if create_three_free([]int{x, y}, player) >= 1 {
			nb_free++
			break
		}
		x--
		y++
	}
	for x, y = coord[0]+1, coord[1]-1; x <= 19 && y >= 0 && Board[x][y] == Board[coord[0]][coord[1]]; {
		if create_three_free([]int{x, y}, player) >= 1 {
			nb_free++
			break
		}
		x++
		y--
	}
	return nb_free
}

func create_three_free(coord []int, player int) int {
	x := coord[0]
	y := coord[1]

	nb_free := 0
	nb_free_hori := 0
	nb_free_verti := 0
	nb_free_slash := 0
	nb_free_anti := 0

	tmp := nb_free
	for i, v, stone := y, 0, 1; (i >= 0) && (Board[x][i] == player || Board[x][i] == NONE) && v <= 3; i-- {
		if Board[x][i] == player && stone < 3 {
			stone++
			nb_free_hori++
		}
		if stone == 3 {
			if i-1 >= 0 && Board[x][i-1] == NONE {
				nb_free++
			}
			break
		}
		v++
	}
	for i, v, stone := y, 0, 1; (i <= 19) && (Board[x][i] == player || Board[x][i] == NONE) && v <= 3; i++ {
		if Board[x][i] == player && stone < 3 {
			stone++
			nb_free_hori++
		}
		if stone == 3 {
			if i+1 <= 19 && Board[x][i+1] == NONE {
				nb_free++
			}
			break
		}
		v++
	}
	if nb_free_hori == 2 && tmp == nb_free {
		nb_free++
	}
	tmp = nb_free
	for i, v, stone := x, 0, 1; (i <= 19) && (Board[i][y] == player || Board[i][y] == NONE) && v <= 3; i++ {
		if Board[i][y] == player && stone < 3 {
			stone++
			nb_free_verti++
		}
		if stone == 3 {
			if i+1 <= 19 && Board[i+1][y] == NONE {
				nb_free++
			}
			break
		}
		v++

	}
	for i, v, stone := x, 0, 1; (i >= 0) && (Board[i][y] == player || Board[i][y] == NONE) && v <= 3; i-- {
		if Board[i][y] == player && stone < 3 {
			stone++
			nb_free_verti++
		}
		if stone == 3 {
			if i-1 >= 0 && Board[i-1][y] == NONE {
				nb_free++
			}
			break
		}
		v++
	}
	if nb_free_verti == 2 && tmp == nb_free {
		nb_free++
	}
	tmp = nb_free
	for x, y, v, stone := coord[0]-1, coord[1]-1, 0, 1; x >= 0 && y >= 0 && (Board[x][y] == player || Board[x][y] == NONE) && v <= 3; {
		if Board[x][y] == player && stone < 3 {
			stone++
			nb_free_slash++
		}
		if stone == 3 {
			if x-1 >= 0 && y-1 >= 0 && Board[x-1][y-1] == NONE {
				nb_free++
			}
			break
		}
		v++
		x--
		y--
	}
	for x, y, v, stone := coord[0]+1, coord[1]+1, 0, 1; x <= 19 && y <= 19 && (Board[x][y] == player || Board[x][y] == NONE) && v <= 3; {
		if Board[x][y] == player && stone < 3 {
			stone++
			nb_free_slash++
		}
		if stone == 3 {
			if x+1 <= 19 && y+1 <= 19 && Board[x+1][y+1] == NONE {
				nb_free++
			}
			break
		}
		v++
		x++
		y++
	}
	if nb_free_slash == 2 && tmp == nb_free {
		nb_free++
	}
	tmp = nb_free
	for x, y, v, stone := coord[0]-1, coord[1]+1, 0, 1; x >= 0 && y <= 19 && (Board[x][y] == player || Board[x][y] == NONE) && v <= 3; {
		if Board[x][y] == player && stone < 3 {
			stone++
			nb_free_anti++
		}
		if stone == 3 {
			if x-1 >= 0 && y+1 <= 19 && Board[x-1][y+1] == NONE {
				nb_free++
			}
			break
		}
		v++
		x--
		y++
	}
	for x, y, v, stone := coord[0]+1, coord[1]-1, 0, 1; x <= 19 && y >= 0 && (Board[x][y] == player || Board[x][y] == NONE) && v <= 3; {
		if Board[x][y] == player && stone < 3 {
			stone++
			nb_free_anti++
		}
		if stone == 3 {
			if x+1 <= 19 && y-1 >= 0 && Board[x+1][y-1] == NONE {
				nb_free++
			}
			break
		}
		v++
		x++
		y--
	}
	if nb_free_anti == 2 && tmp == nb_free {
		nb_free++
	}
	return nb_free
}

func dual_three(coord []int, player int) bool {

	res := create_three_free(coord, player)
	if res > 1 {
		fmt.Print("\nDUAL THREE\n")
		return true
	}
	if res == 1 {
		fmt.Print("\nUN THREE\n")
		ff := check_all_three_free(coord, player)
		if ff >= 6 {
			fmt.Printf("\nencore '%d' THREE\n", ff)
			return true
		}
	}
	return false
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
		capt[i] = []int{x + 1, y + 1}
		capt[i+1] = []int{x + 2, y + 2}
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
		if breakable([]int{x, i}) == false {
			hor++
		} else {
			break
		}
	}
	for i := y; (i <= 19) && Board[x][i] == Board[x][y]; i++ {
		if breakable([]int{x, i}) == false {
			hor++
		} else {
			break
		}
	}
	for i := x; (i <= 19) && Board[i][y] == Board[x][y]; i++ {
		if breakable([]int{i, y}) == false {
			vert++
		} else {
			break
		}
	}
	for i := x; (i >= 0) && Board[i][y] == Board[x][y]; i-- {
		if breakable([]int{i, y}) == false {
			vert++
		} else {
			break
		}
	}
	// check \
	for x, y = coord[0]-1, coord[1]-1; x >= 0 && y >= 0 && Board[x][y] == Board[coord[0]][coord[1]]; {
		if breakable([]int{x, y}) == false {
			tl++
		} else {
			break
		}
		x--
		y--
	}
	for x, y = coord[0]+1, coord[1]+1; x <= 19 && y <= 19 && Board[x][y] == Board[coord[0]][coord[1]]; {
		if breakable([]int{x, y}) == false {
			tl++
		} else {
			break
		}
		x++
		y++
	}

	//check /
	for x, y = coord[0]-1, coord[1]+1; x >= 0 && y <= 19 && Board[x][y] == Board[coord[0]][coord[1]]; {
		if breakable([]int{x, y}) == false {
			tr++
		} else {
			break
		}
		x--
		y++
	}
	for x, y = coord[0]+1, coord[1]-1; x <= 19 && y >= 0 && Board[x][y] == Board[coord[0]][coord[1]]; {
		if breakable([]int{x, y}) == false {
			tr++
		} else {
			break
		}
		x++
		y--
	}

	if hor > NB_ALIGN_TO_WIN || vert > NB_ALIGN_TO_WIN || tl >= NB_ALIGN_TO_WIN || tr >= NB_ALIGN_TO_WIN {
		return true
	}
	return false
}

func check_win(coord []int) (bool, int) {
	if check_align(coord) == true {
		return true, getinvturn()
	}
	if BPOW == 10 {
		return true, BLACK
	}
	if WPOW == 10 {
		return true, WHITE
	}
	return false, -1
}

func referee(coord []string, ws *websocket.Conn) (bool, bool, int) {

	x, _ := strconv.Atoi(coord[0])
	y, _ := strconv.Atoi(coord[1])

	coordint := []int{x, y}

	if Board[x][y] == NONE {
		if dual_three(coordint, Turn) == false {
			Board[x][y] = Turn
			if Turn == BLACK {
				Turn = WHITE
			} else {
				Turn = BLACK
			}
			send_capture(capture(coordint), ws)
			win, who := check_win(coordint)
			return true, win, who
		}
	}
	return false, false, -1
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

func getinvturn() int {
	if Turn == BLACK {
		return WHITE
	}
	return BLACK
}

func getStringTurn() string {
	if Turn == BLACK {
		return "black"
	}
	return "white"
}

func getStringTurnInv() string {
	if Turn == BLACK {
		return "white"
	}
	return "black"
}

func getStringPl(pl int) string {
	if pl == BLACK {
		return "black"
	}
	return "white"
}

func sendRecvCoord(ws *websocket.Conn) {
	for {
		var buf string

		buf = ws_recv(ws)

		//check avec le referee
		if buf == "reset" {
			Board = initBoard(GOBANSIZE)

		} else if buf == "getturn" {
			ws_send("turn,"+getStringTurn(), ws)
		} else if buf == "getscore" {
			ws_send("score, Black : "+strconv.Itoa(BPOW)+" | White : "+strconv.Itoa(WPOW), ws)
		} else {
			mov, win, who := referee(strings.Split(buf, ","), ws)
			if win {
				buf = "win," + getStringPl(who)
			} else if mov == true {
				buf += "," + getStringTurnInv()
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
