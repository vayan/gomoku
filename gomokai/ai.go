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

type Tree struct {
	note     int
	coo      Coord
	branches []Tree
}

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

func choose_move_tree(tree Tree, test int) {
	for _, branche := range tree.branches {
		choose_move_tree(branche, test+1)
		log.Printf("Depth %d %d", test, branche.note)
	}
}

func choose_move() Coord {
	log.Printf("Note : Black %d White %d", movehim.note, moveme.note)
	if movehim.note >= moveme.note {
		return movehim.coo
	}
	return moveme.coo
}

func minimax(board [][]int, turn int, depth int, or_depth int) Tree {
	var (
		mtree Tree
		xtree Tree
		ytree Tree
	)
	mtree = Tree{0, Coord{0, 0}, make([]Tree, 22)}
	for x := 0; x < 20; x++ {
		xtree = Tree{0, Coord{x, 0}, make([]Tree, 22)}
		for y := 0; y < 20; y++ {
			ytree = Tree{0, Coord{x, y}, make([]Tree, 22)}
			// TODO stock les move interdit en concu et compare apres
			// TODO evite de scan toute les cases 50%done 
			if can_move(Coord{x, y}, board, turn, depth, or_depth) {
				if depth > 0 {
					next_board := duplicate_board(board)
					next_board[x][y] = turn
					child_tree := minimax(next_board, get_opos_turn(turn), depth-1, or_depth)
					ytree.branches[y] = child_tree //test lenteur
				}
				note := eval_move(Coord{x, y}, board, turn)
				ytree.note = note
				if turn == BLACK {
					ytree.note = note * -1
				}
				if turn == BLACK {
					if movehim.note < note {
						movehim = Move{Coord{x, y}, note}
					}
				} else if moveme.note < note {
					moveme = Move{Coord{x, y}, note}
				}
				xtree.branches[x] = ytree
			}
		}
		mtree.branches[x] = xtree
	}
	return mtree
}

func get_heat_areas() []Area {
	return nil
}

func start_ai() string {
	depth := 1
	movehim = Move{Coord{-1, -1}, 0}
	moveme = Move{Coord{-1, -1}, 0}

	log.Print("== AI Start thinking ==")
	t0 := time.Now()
	tree := minimax(duplicate_board(Board), Turn, depth, depth)
	mov := choose_move()
	log.Print("Start choosing from tree")
	//choose_move_tree(tree, 0)
	t1 := time.Now()
	log.Printf("The AI took %v to think.\n", t1.Sub(t0))
	return "PLAY " + strconv.Itoa(mov.x) + " " + strconv.Itoa(mov.y)
}
