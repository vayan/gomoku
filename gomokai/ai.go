package main

import (
	"log"
	"strconv"
)

type Move struct {
	coo  Coord
	note int
	next *map[Move]int
}

type Coord struct {
	x int
	y int
}

type Area struct {
	topleft     Coord
	bottomright Coord
}

func find_ennemy() {

}

func choose_move(moves map[Move]int) Coord {

	var note = 0
	var choosen_one = Coord{-1, -1}

	for mov, _ := range moves {
		if note < mov.note {
			note = mov.note
			choosen_one = mov.coo
		}
	}
	return choosen_one
}

func all_move(board [][]int, turn int, depth int, move_tree map[Move]int) int {

	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			if can_move(Coord{x, y}, board, turn) {
				var Moves = make(map[Move]int)

				move := Move{Coord{x, y}, 0, &Moves}
				if depth > 0 {
					next_board := duplicate_board()
					next_board[x][y] = turn
					minimax(depth-1, get_opos_turn(turn), next_board, Moves)
				}
				move.note = eval_move(Coord{x, y}, board, turn)
				move_tree[move] = 0
			}
		}
	}
	return depth
}

func get_heat_areas() []Area {
	return nil
}

func minimax(depth int, turn int, origin [][]int, move_tree map[Move]int) {

	if depth > 0 {
		new_depth := all_move(origin, turn, depth, move_tree)
		depth = new_depth
	}
}

func start_ai() string {
	log.Print("== AI Start thinking ==")
	tree := make(map[Move]int)
	minimax(1, Turn, duplicate_board(), tree)
	mov := choose_move(tree)
	return "PLAY " + strconv.Itoa(mov.x) + " " + strconv.Itoa(mov.y)
}
