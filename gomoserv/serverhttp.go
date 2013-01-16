package main

import (
	"html/template"
	"net/http"
)

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
