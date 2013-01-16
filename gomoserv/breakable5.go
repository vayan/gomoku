package main

func breakable_opos(coord []int) bool {
	x := coord[0]
	y := coord[1]
	to_capt := BLACK

	if Board[x][y] == BLACK {
		to_capt = WHITE
	}

	if (x-1 >= 0 && x+2 <= 19 && Board[x-1][y] == to_capt && Board[x+1][y] == Board[x][y] && Board[x+2][y] == NONE) ||
		(x-2 >= 0 && x+2 <= 19 && Board[x-1][y] == Board[x][y] && Board[x+1][y] == to_capt && Board[x-2][y] == NONE) ||
		(y+1 <= 19 && y-2 >= 0 && Board[x][y+1] == to_capt && Board[x][y-1] == Board[x][y] && Board[x][y-2] == NONE) ||
		(y+2 <= 19 && y-1 >= 0 && Board[x][y+1] == Board[x][y] && Board[x][y-1] == to_capt && Board[x][y+2] == NONE) ||
		(x-2 >= 0 && y+2 <= 19 && x+2 <= 19 && y-2 >= 0 && Board[x-1][y+1] == Board[x][y] && Board[x+1][y-1] == to_capt && Board[x-2][y+2] == NONE) ||
		(x-2 >= 0 && y+2 <= 19 && x+2 <= 19 && y-2 >= 0 && Board[x-1][y+1] == to_capt && Board[x+1][y-1] == Board[x][y] && Board[x+2][y-2] == NONE) ||
		(x-2 >= 0 && y-2 >= 0 && x+2 <= 19 && y+2 <= 19 && Board[x-1][y-1] == Board[x][y] && Board[x+1][y+1] == to_capt && Board[x-2][y-2] == NONE) ||
		(x-2 >= 0 && y-2 >= 0 && x+2 <= 19 && y+2 <= 19 && Board[x-1][y-1] == to_capt && Board[x+1][y+1] == Board[x][y] && Board[x+2][y+2] == NONE) {
		return true
	}
	return false
}

func breakable_same(coord []int) bool {
	x := coord[0]
	y := coord[1]
	to_capt := BLACK

	if Board[x][y] == BLACK {
		to_capt = WHITE
	}

	if (x-2 >= 0 && x+1 <= 19 && Board[x-1][y] == Board[x][y] && Board[x-2][y] == to_capt && Board[x+1][y] == NONE) ||
		(x+2 <= 19 && x-1 >= 0 && Board[x+1][y] == Board[x][y] && Board[x+2][y] == to_capt && Board[x-1][y] == NONE) ||
		(y+2 <= 19 && y-1 >= 0 && Board[x][y+1] == Board[x][y] && Board[x][y+2] == to_capt && Board[x][y-1] == NONE) ||
		(y-2 >= 0 && y+1 <= 19 && Board[x][y-1] == Board[x][y] && Board[x][y-2] == to_capt && Board[x][y+1] == NONE) ||
		(x-2 >= 0 && y+2 <= 19 && x+1 <= 19 && y-1 >= 0 && Board[x-1][y+1] == Board[x][y] && Board[x-2][y+2] == to_capt && Board[x+1][y-1] == NONE) ||
		(x+2 <= 19 && y-2 >= 0 && x-1 >= 0 && y+1 <= 19 && Board[x+1][y-1] == Board[x][y] && Board[x+2][y-2] == to_capt && Board[x-1][y+1] == NONE) ||
		(x-2 >= 0 && y-2 >= 0 && x+1 <= 19 && y+1 <= 19 && Board[x-1][y-1] == Board[x][y] && Board[x-2][y-2] == to_capt && Board[x+1][y+1] == NONE) ||
		(x+2 <= 19 && y+2 <= 19 && x-1 >= 0 && y-1 >= 0 && Board[x+1][y+1] == Board[x][y] && Board[x+2][y+2] == to_capt && Board[x-1][y-1] == NONE) {
		return true
	}
	return false
}

func breakable(coord []int) bool {
	if BREAKING_5 == 0 {
		return false
	}
	if breakable_same(coord) {
		return true
	} else if breakable_opos(coord) {
		return true
	}
	return false
}
