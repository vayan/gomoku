package main

import ()

func connected(move Coord, board [][]int, type_conn int) (int, int, int, int) {
	var (
		x     = move.x
		y     = move.y
		hori  = 0
		verti = 0
		diagl = 0
		diagr = 0
	)

	//CHECK ----
	for i := y; (i >= 0) && board[x][i] == type_conn; i-- {
		hori++
	}
	for i := y; (i <= 19) && board[x][i] == type_conn; i++ {
		hori++
	}

	//CHECK |
	for i := x; (i <= 19) && board[i][y] == type_conn; i++ {
		verti++
	}
	for i := x; (i >= 0) && board[i][y] == type_conn; i-- {
		verti++
	}

	//CHECK \
	for x, y = move.x, move.y; x >= 0 && y >= 0 && board[x][y] == type_conn; {
		diagl++
		x--
		y--
	}

	for x, y = move.x, move.y; x <= 19 && y <= 19 && board[x][y] == type_conn; {
		diagl++
		x++
		y++
	}

	//check /
	for x, y = move.x, move.y; x >= 0 && y <= 19 && Board[x][y] == type_conn; {
		diagr++
		x--
		y++
	}
	for x, y = move.x, move.y; x <= 19 && y >= 0 && Board[x][y] == type_conn; {
		diagr++
		x++
		y--
	}

	return hori, verti, diagl, diagr
}

func can_move(move Coord, board [][]int, turn int) bool {
	//log.Print("Test if can put stone here")
	if board[move.x][move.y] != NONE {
		//log.Print("Already stone here")
		return false
	}
	if breakable(move, board) {
		return false
	}
	if dual_three(move, turn, board) {
		return false
	}

	return true
}

func eval_connection(move Coord, board [][]int, turn int) int {
	//check connect friend
	fhori, fverti, fdiagl, fdiagr := connected(move, board, turn)

	return fhori + fverti + fdiagl + fdiagr
}

func eval_move(move Coord, board [][]int, turn int) int {
	board[move.x][move.y] = turn //on pose la pierre temp
	ret := eval_connection(move, board, turn)
	board[move.x][move.y] = NONE //restore l'etat
	return ret
}
