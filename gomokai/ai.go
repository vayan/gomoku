package main

import (
	"fmt"
	"strconv"
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

func choose_move_in_tree(tree Tree) Coord {
	var (
		choos = Coord{0, 0}
		not   = 0
	)
	for _, branche := range tree.branches {
		if branche.note > not {
			choos = branche.coo
		}
	}
	return choos
}

func walk_in_tree(tree Tree, coo Coord, not int) (Tree, Coord, int) {
	for _, branche := range tree.branches {
		if len(branche.branches) > 0 {
			_, coo, not = walk_in_tree(branche, coo, not)
			if branche.note > not {
				not = branche.note
				coo = branche.coo
			}
			tree.note += branche.note
			//log.Print(" ", branche.note)
		}

	}

	return tree, coo, not
}

func choose_move() Coord {
	//log.Printf("Note : Black %d White %d", movehim.note, moveme.note)
	if movehim.note >= moveme.note {
		return movehim.coo
	}
	return moveme.coo
}

func minimax(board [][]int, turn int, depth int, or_depth int, newmov Coord, speed int) Tree {
	var (
		mtree      Tree
		child_tree Tree
		val        = 0
	)
	_ = mtree
	_ = child_tree
	if speed > 0 {
		mtree = Tree{0, Coord{newmov.x, newmov.y}, make([]Tree, 20*20)}
	}
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			// TODO stock les move interdit en concu et compare apres
			// TODO evite de scan toute les cases 50%done 
			if can_move(Coord{x, y}, board, turn, depth, or_depth) {
				if depth > 0 {
					next_board := duplicate_board(board)
					next_board[x][y] = turn
					child_tree = minimax(next_board, get_opos_turn(turn), depth-1, or_depth, Coord{x, y}, speed)
				} else if speed > 0 {
					child_tree = Tree{0, Coord{x, y}, make([]Tree, 20*20)}
				}
				note := eval_move(Coord{x, y}, board, turn)
				if speed > 0 {
					child_tree.note = note
					if turn == BLACK {
						child_tree.note = note + 100
					}
					mtree.branches[val] = child_tree
				}
				if turn == BLACK {
					if movehim.note < note {
						movehim = Move{Coord{x, y}, note}
					}
				} else if moveme.note < note {
					moveme = Move{Coord{x, y}, note}
				}
				val++
			}

		}
	}
	return mtree
}

func get_heat_areas() []Area {
	return nil
}

func aff_score(tree Tree) {
	for _, branche := range tree.branches {
		fmt.Printf("%d ", branche.note)
	}
	fmt.Print("\n")
}

func start_ai() string {
	mov := Coord{0, 0}
	depth := 1
	speed := 0
	movehim = Move{Coord{-1, -1}, 0}
	moveme = Move{Coord{-1, -1}, 0}

	//log.Print("== AI Start thinking ==")

	//log.Print("AI trying all moves")
	//t0 := time.Now()
	tree := minimax(duplicate_board(Board), Turn, depth, depth, Coord{0, 0}, speed)
	//t1 := time.Now()
	//log.Printf("The AI took %v to check all move.....", t1.Sub(t0))

	//log.Print("Start choosing from tree")
	//t0 = time.Now()
	if speed > 1 {
		_, mov, _ = walk_in_tree(tree, Coord{0, 0}, 0)
	} else {
		mov = choose_move()
	}
	//t1 = time.Now()
	//log.Printf("The AI took %v to choose one.....", t1.Sub(t0))

	//aff_score(tree)
	return "PLAY " + strconv.Itoa(mov.x) + " " + strconv.Itoa(mov.y)
}
