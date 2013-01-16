package main

import (
	"log"
	"strconv"
	"time"
)

var (
	movehim = Move{Coord{-1, -1}, 0}
	moveme  = Move{Coord{-1, -1}, 0}
)

type Move struct {
	coo  Coord
	note int
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

func choose_move() Coord {
	log.Printf("Note : Black %d White %d", movehim.note, moveme.note)
	if movehim.note >= moveme.note {
		return movehim.coo
	}
	return moveme.coo
}

func all_move(board [][]int, turn int, depth int) int {

	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			//stock les move interdit en parrale et compare apres
			if can_move(Coord{x, y}, board, turn) {
				if depth > 0 {
					next_board := duplicate_board(board)
					next_board[x][y] = turn
					all_move(next_board, get_opos_turn(turn), depth-1)
				}
				note := eval_move(Coord{x, y}, board, turn)
				if turn == BLACK {
					if movehim.note < note {
						movehim = Move{Coord{x, y}, note}
					}
				} else if moveme.note < note {
					moveme = Move{Coord{x, y}, note}
				}
			}
		}
	}
	return depth
}

func get_heat_areas() []Area {
	return nil
}

func minimax(depth int, turn int, origin [][]int) {

	all_move(origin, turn, depth)
}

func start_ai() string {
	movehim = Move{Coord{-1, -1}, 0}
	moveme = Move{Coord{-1, -1}, 0}
	log.Print("== AI Start thinking ==")
	t0 := time.Now()
	minimax(1, Turn, duplicate_board(Board))
	mov := choose_move()
	t1 := time.Now()
	log.Printf("The AI took %v to think.\n", t1.Sub(t0))
	return "PLAY " + strconv.Itoa(mov.x) + " " + strconv.Itoa(mov.y)
}
