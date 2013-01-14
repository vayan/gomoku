package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Add(buff []string) {
	x, _ := strconv.Atoi(buff[1])
	y, _ := strconv.Atoi(buff[2])
	addStone(x, y, Turn)
	change_turn()
	AffBoard(20)
}

func Rm(buff []string) {
	x1, _ := strconv.Atoi(buff[1])
	y1, _ := strconv.Atoi(buff[2])
	rmStone(x1, y1)
	x2, _ := strconv.Atoi(buff[3])
	y2, _ := strconv.Atoi(buff[4])
	rmStone(x2, y2)
	AffBoard(20)
}

func parser(msg string) {
	buff := strings.Split(msg, " ")
	buf := buff[0]

	switch buf {
	case "ADD":
		Add(buff)
	case "REM":
		Rm(buff)
	case "WIN":
		log.Print("AI WIN")
		os.Exit(11)
	case "LOSE":
		log.Print("AI LOSE")
		os.Exit(11)
	case "YOURTURN":
		start_ai()
	case "RULES":
	}
}
