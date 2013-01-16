package main

import (
	"strconv"
)

func capture(coord []int) [][]int {
	x := coord[0]
	y := coord[1]
	i := 0

	var capt = make([][]int, 19)
	to_capt := BLACK
	if Board[x][y] == BLACK {
		to_capt = WHITE
	}

	// hori

	if x-3 >= 0 && Board[x-1][y] == to_capt && Board[x-2][y] == to_capt && Board[x-3][y] == Board[x][y] {
		capt[i] = []int{x - 1, y}
		capt[i+1] = []int{x - 2, y}
		i += 2
	}
	if x+3 <= 19 && Board[x+1][y] == to_capt && Board[x+2][y] == to_capt && Board[x+3][y] == Board[x][y] {
		capt[i] = []int{x + 1, y}
		capt[i+1] = []int{x + 2, y}
		i += 2
	}

	//verti
	if y+3 <= 19 && Board[x][y+1] == to_capt && Board[x][y+2] == to_capt && Board[x][y+3] == Board[x][y] {
		capt[i] = []int{x, y + 1}
		capt[i+1] = []int{x, y + 2}
		i += 2
	}
	if y-3 >= 0 && Board[x][y-1] == to_capt && Board[x][y-2] == to_capt && Board[x][y-3] == Board[x][y] {
		capt[i] = []int{x, y - 1}
		capt[i+1] = []int{x, y - 2}
		i += 2
	}

	// /
	if x-3 >= 0 && y+3 <= 19 && Board[x-1][y+1] == to_capt && Board[x-2][y+2] == to_capt && Board[x-3][y+3] == Board[x][y] {
		capt[i] = []int{x - 1, y + 1}
		capt[i+1] = []int{x - 2, y + 2}
		i += 2
	}
	if x+3 <= 19 && y-3 >= 0 && Board[x+1][y-1] == to_capt && Board[x+2][y-2] == to_capt && Board[x+3][y-3] == Board[x][y] {
		capt[i] = []int{x + 1, y - 1}
		capt[i+1] = []int{x + 2, y - 2}
		i += 2
	}

	// \
	if x-3 >= 0 && y-3 >= 0 && Board[x-1][y-1] == to_capt && Board[x-2][y-2] == to_capt && Board[x-3][y-3] == Board[x][y] {
		capt[i] = []int{x - 1, y - 1}
		capt[i+1] = []int{x - 2, y - 2}
		i += 2
	}
	if x+3 <= 19 && y+3 <= 19 && Board[x+1][y+1] == to_capt && Board[x+2][y+2] == to_capt && Board[x+3][y+3] == Board[x][y] {
		capt[i] = []int{x + 1, y + 1}
		capt[i+1] = []int{x + 2, y + 2}
		i += 2
	}
	if i != 0 {
		return capt
	}
	return nil
}

func check_align(coord []int) bool {
	hor := 0
	vert := 0
	tl := 1
	tr := 1
	x := coord[0]
	y := coord[1]

	//check hor
	for i := y; (i >= 0) && Board[x][i] == Board[x][y]; i-- {
		if breakable([]int{x, i}) == false {
			hor++
		} else {
			break
		}
	}
	for i := y; (i <= 19) && Board[x][i] == Board[x][y]; i++ {
		if breakable([]int{x, i}) == false {
			hor++
		} else {
			break
		}
	}
	for i := x; (i <= 19) && Board[i][y] == Board[x][y]; i++ {
		if breakable([]int{i, y}) == false {
			vert++
		} else {
			break
		}
	}
	for i := x; (i >= 0) && Board[i][y] == Board[x][y]; i-- {
		if breakable([]int{i, y}) == false {
			vert++
		} else {
			break
		}
	}
	// check \
	for x, y = coord[0]-1, coord[1]-1; x >= 0 && y >= 0 && Board[x][y] == Board[coord[0]][coord[1]]; {
		if breakable([]int{x, y}) == false {
			tl++
		} else {
			break
		}
		x--
		y--
	}
	for x, y = coord[0]+1, coord[1]+1; x <= 19 && y <= 19 && Board[x][y] == Board[coord[0]][coord[1]]; {
		if breakable([]int{x, y}) == false {
			tl++
		} else {
			break
		}
		x++
		y++
	}

	//check /
	for x, y = coord[0]-1, coord[1]+1; x >= 0 && y <= 19 && Board[x][y] == Board[coord[0]][coord[1]]; {
		if breakable([]int{x, y}) == false {
			tr++
		} else {
			break
		}
		x--
		y++
	}
	for x, y = coord[0]+1, coord[1]-1; x <= 19 && y >= 0 && Board[x][y] == Board[coord[0]][coord[1]]; {
		if breakable([]int{x, y}) == false {
			tr++
		} else {
			break
		}
		x++
		y--
	}

	if hor > NB_ALIGN_TO_WIN || vert > NB_ALIGN_TO_WIN || tl >= NB_ALIGN_TO_WIN || tr >= NB_ALIGN_TO_WIN {
		return true
	}
	return false
}

func check_win(coord []int) (bool, int) {
	if check_align(coord) == true {
		return true, getinvturn()
	}
	if BPOW == 10 {
		return true, BLACK
	}
	if WPOW == 10 {
		return true, WHITE
	}
	return false, -1
}

func referee(coord []string, c Connection) (bool, bool, int) {

	for pl, _ := range players {
		if (pl == c) && pl.player_color != Turn {
			return false, false, -1
		}
	}

	x, _ := strconv.Atoi(coord[0])
	y, _ := strconv.Atoi(coord[1])

	coordint := []int{x, y}

	if Board[x][y] == NONE {
		if dual_three(coordint, Turn) == false {
			Board[x][y] = Turn
			if Turn == BLACK {
				Turn = WHITE
			} else {
				Turn = BLACK
			}
			send_capture(capture(coordint), c)
			win, who := check_win(coordint)
			return true, win, who
		}
	}
	return false, false, -1
}
