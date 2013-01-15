package main

import (
	"log"
)

type Scan struct {
	hori  string
	verti string
	diag  string
}

/*
   C = cur move
   E = Ennemy
   M = ME
   S = Empty
*/

func can_move(move Coord, board [][]int, turn int) bool {
	//log.Print("Test if can put stone here")
	if board[move.x][move.y] != NONE {
		log.Print("Already stone here")
		return false
	}
	return true
}

func eval_move(move Coord, board [][]int, turn int) int {
	return randInt(1, 90)
}
