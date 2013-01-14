package main

import (
	"html/template"
	"net/http"
	"strconv"
)

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
				board += " <td><div class='" + color + " stone pos" + strconv.Itoa(x) + "y" + strconv.Itoa(y) + "'>" + strconv.Itoa(x) + " " + strconv.Itoa(y) + "</div></td> "
			} else {
				board += " <td></td> "
			}
		}
		board += "</tr>"
	}
	return board
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
