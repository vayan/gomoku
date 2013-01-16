package main

func breakable_opos(coord Coord, board [][]int) bool {
	x := coord.x
	y := coord.y
	to_capt := BLACK

	if board[x][y] == BLACK {
		to_capt = WHITE
	}

	if (x-1 >= 0 && x+2 <= 19 && board[x-1][y] == to_capt && board[x+1][y] == board[x][y] && board[x+2][y] == NONE) ||
		(x-2 >= 0 && x+2 <= 19 && board[x-1][y] == board[x][y] && board[x+1][y] == to_capt && board[x-2][y] == NONE) ||
		(y+1 <= 19 && y-2 >= 0 && board[x][y+1] == to_capt && board[x][y-1] == board[x][y] && board[x][y-2] == NONE) ||
		(y+2 <= 19 && y-1 >= 0 && board[x][y+1] == board[x][y] && board[x][y-1] == to_capt && board[x][y+2] == NONE) ||
		(x-2 >= 0 && y+2 <= 19 && x+2 <= 19 && y-2 >= 0 && board[x-1][y+1] == board[x][y] && board[x+1][y-1] == to_capt && board[x-2][y+2] == NONE) ||
		(x-2 >= 0 && y+2 <= 19 && x+2 <= 19 && y-2 >= 0 && board[x-1][y+1] == to_capt && board[x+1][y-1] == board[x][y] && board[x+2][y-2] == NONE) ||
		(x-2 >= 0 && y-2 >= 0 && x+2 <= 19 && y+2 <= 19 && board[x-1][y-1] == board[x][y] && board[x+1][y+1] == to_capt && board[x-2][y-2] == NONE) ||
		(x-2 >= 0 && y-2 >= 0 && x+2 <= 19 && y+2 <= 19 && board[x-1][y-1] == to_capt && board[x+1][y+1] == board[x][y] && board[x+2][y+2] == NONE) {
		return true
	}
	return false
}

func breakable_same(coord Coord, board [][]int) bool {
	x := coord.x
	y := coord.y
	to_capt := BLACK

	if board[x][y] == BLACK {
		to_capt = WHITE
	}

	if (x-2 >= 0 && x+1 <= 19 && board[x-1][y] == board[x][y] && board[x-2][y] == to_capt && board[x+1][y] == NONE) ||
		(x+2 <= 19 && x-1 >= 0 && board[x+1][y] == board[x][y] && board[x+2][y] == to_capt && board[x-1][y] == NONE) ||
		(y+2 <= 19 && y-1 >= 0 && board[x][y+1] == board[x][y] && board[x][y+2] == to_capt && board[x][y-1] == NONE) ||
		(y-2 >= 0 && y+1 <= 19 && board[x][y-1] == board[x][y] && board[x][y-2] == to_capt && board[x][y+1] == NONE) ||
		(x-2 >= 0 && y+2 <= 19 && x+1 <= 19 && y-1 >= 0 && board[x-1][y+1] == board[x][y] && board[x-2][y+2] == to_capt && board[x+1][y-1] == NONE) ||
		(x+2 <= 19 && y-2 >= 0 && x-1 >= 0 && y+1 <= 19 && board[x+1][y-1] == board[x][y] && board[x+2][y-2] == to_capt && board[x-1][y+1] == NONE) ||
		(x-2 >= 0 && y-2 >= 0 && x+1 <= 19 && y+1 <= 19 && board[x-1][y-1] == board[x][y] && board[x-2][y-2] == to_capt && board[x+1][y+1] == NONE) ||
		(x+2 <= 19 && y+2 <= 19 && x-1 >= 0 && y-1 >= 0 && board[x+1][y+1] == board[x][y] && board[x+2][y+2] == to_capt && board[x-1][y-1] == NONE) {
		return true
	}
	return false
}

func breakable(coord Coord, board [][]int) bool {
	if BREAKING_5 == 0 {
		return false
	}
	if breakable_same(coord, board) {
		return true
	} else if breakable_opos(coord, board) {
		return true
	}
	return false
}
