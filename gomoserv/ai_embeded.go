package main

import (
	"gomoku/gomokai"
	"strconv"
)

func hint(con Connection) {
	_, x, y := gomokai.Start_ai(Board)
	send("HINT "+strconv.Itoa(x)+" "+strconv.Itoa(y), con)
}
