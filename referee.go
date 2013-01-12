package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"strconv"
)

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
	if breakable_same(coord) {
		return true
	} else if breakable_opos(coord) {
		return true
	}
	return false
}

func check_three_free(coord []int, player int, sens int) (int, int) {
	x := coord[0]
	y := coord[1]

	xori := coord[0]
	yori := coord[1]

	nbr := 0
	nb_free_hori := 0
	nb_free_verti := 0
	nb_free_slash := 0
	nb_free_anti := 0
	tmp := nbr
	limit := 4

	if sens != 1 && sens != 2 {
		for i, v, stone := y, 0, 0; (i >= 0) && (Board[x][i] == player || Board[x][i] == NONE) && v <= 3; i-- {
			if Board[x][i] == player {
				stone++
				nb_free_hori++
			}
			if stone == limit && i-1 >= 0 && yori+1 <= 19 && Board[x][yori+1] == NONE && Board[x][i-1] == NONE {
				fmt.Print("Find three y--\n")
				nbr++
				return 1, 1
			}
			v++
		}
		for i, v, stone := y, 0, 0; (i <= 19) && (Board[x][i] == player || Board[x][i] == NONE) && v <= 3; i++ {
			if Board[x][i] == player {
				stone++
				nb_free_hori++
			}
			if stone == limit && i+1 <= 19 && yori-1 >= 0 && Board[x][yori-1] == NONE && Board[x][i+1] == NONE {
				fmt.Print("Find three y++\n")
				nbr++
				return 2, 1
			}
			v++
		}
		if nb_free_hori == limit && tmp == nbr {
			fmt.Print("Find three _\n")
			return 2, 1
		}
	}

	if sens != 3 && sens != 4 {
		tmp = nbr
		for i, v, stone := x, 0, 0; (i <= 19) && (Board[i][y] == player || Board[i][y] == NONE) && v <= 3; i++ {
			if Board[i][y] == player {
				stone++
				nb_free_verti++
			}
			if stone == limit && i+1 <= 19 && xori-1 >= 0 && Board[xori-1][y] == NONE && Board[i+1][y] == NONE {
				fmt.Print("Find three x++\n")
				nbr++
				return 3, 1
			}
			v++

		}
		for i, v, stone := x, 0, 0; (i >= 0) && (Board[i][y] == player || Board[i][y] == NONE) && v <= 3; i-- {
			if Board[i][y] == player {
				stone++
				nb_free_verti++
			}
			if stone == limit && i-1 >= 0 && xori+1 <= 19 && Board[xori+1][y] == NONE && Board[i-1][y] == NONE {
				fmt.Print("Find three x--\n")
				nbr++
				return 4, 1
			}
			v++
		}
		if nb_free_verti == limit && tmp == nbr {
			fmt.Print("Find three |\n")
			return 4, 1
		}
	}
	if sens != 5 && sens != 6 {
		tmp = nbr
		for x, y, v, stone := coord[0], coord[1], 0, 0; x >= 0 && y >= 0 && (Board[x][y] == player || Board[x][y] == NONE) && v <= 3; {
			if Board[x][y] == player {
				stone++
				nb_free_slash++
			}
			if stone == limit && x-1 >= 0 && y-1 >= 0 && xori+1 <= 19 && yori+1 <= 19 && Board[xori+1][yori+1] == NONE && Board[x-1][y-1] == NONE {
				fmt.Print("Find three x- y-\n")
				nbr++
				return 5, 1
			}
			v++
			x--
			y--
		}
		for x, y, v, stone := coord[0], coord[1], 0, 0; x <= 19 && y <= 19 && (Board[x][y] == player || Board[x][y] == NONE) && v <= 3; {
			if Board[x][y] == player {
				stone++
				nb_free_slash++
			}
			if stone == limit && x+1 <= 19 && y+1 <= 19 && xori-1 >= 0 && yori-1 >= 0 && Board[xori-1][yori-1] == NONE && Board[x+1][y+1] == NONE {
				fmt.Print("Find three x+ y+\n")
				nbr++
				return 6, 1
			}
			v++
			x++
			y++
		}
		if nb_free_slash == limit && tmp == nbr {
			fmt.Print("Find three \\ \n")
			return 6, 1
		}
	}
	if sens != 7 && sens != 8 {
		tmp = nbr
		for x, y, v, stone := coord[0], coord[1], 0, 0; x >= 0 && y <= 19 && (Board[x][y] == player || Board[x][y] == NONE) && v <= 3; {
			if Board[x][y] == player {
				stone++
				nb_free_anti++
			}
			if stone == limit && x-1 >= 0 && y+1 <= 19 && xori+1 <= 19 && yori-1 >= 0 && Board[xori+1][yori-1] == NONE && Board[x-1][y+1] == NONE {
				fmt.Print("Find three x- y+\n")
				nbr++
				return 7, 1
			}
			v++
			x--
			y++
		}

		for x, y, v, stone := coord[0], coord[1], 0, 0; x <= 19 && y >= 0 && (Board[x][y] == player || Board[x][y] == NONE) && v <= 3; {
			if Board[x][y] == player {
				stone++
				nb_free_anti++
			}
			if stone == limit && x+1 <= 19 && y-1 >= 0 && xori-1 >= 0 && yori+1 <= 19 && Board[xori-1][yori+1] == NONE && Board[x+1][y-1] == NONE {
				fmt.Print("Find three x+ y-\n")
				nbr++
				return 8, 1
			}
			v++
			x++
			y--
		}
		if nb_free_anti == limit && tmp == nbr {
			fmt.Print("Find three / \n")
			return 8, 1
		}
	}
	return 0, 0
}

func check_dual_three(coord []int, player int, sens int) bool {
	x := coord[0]
	y := coord[1]

	xori := coord[0]
	yori := coord[1]

	limit := 0

	nbr := 0

	var BoardVisited [][]int = make([][]int, 20)

	for xx := 0; xx < 20; xx++ {
		BoardVisited[xx] = make([]int, 20)
		for yy := 0; yy < 20; yy++ {
			BoardVisited[xx][yy] = 0
		}
	}
	BoardVisited[x][y] = 1
	if sens != 1 && sens != 2 {
		fmt.Print("=====Check les sens _\n")
		for i := y; (i >= 0) && Board[x][i] == Board[x][y]; i-- {
			if BoardVisited[x][i] == 0 {
				fmt.Print("==Check first y-\n")
				nb, _ := check_three_free([]int{x, i}, player, sens)
				if nb > limit {
					nbr++
				}
				BoardVisited[x][i] = 1
			}
		}
		for i := y; (i <= 19) && Board[x][i] == Board[x][y]; i++ {
			if BoardVisited[x][i] == 0 {
				fmt.Print("==Check first y+\n")
				nb, _ := check_three_free([]int{x, i}, player, sens)
				if nb > limit {
					nbr++
				}
				BoardVisited[x][i] = 1
			}
		}
	}
	if sens != 3 && sens != 4 {
		fmt.Print("=====Check les sens |\n")
		for i := x; (i <= 19) && Board[i][y] == Board[x][y]; i++ {
			if BoardVisited[i][y] == 0 {
				fmt.Print("==Check first x+\n")
				nb, _ := check_three_free([]int{i, y}, player, sens)
				if nb > limit {
					nbr++
				}
				BoardVisited[i][y] = 1
			}
		}
		for i := x; (i >= 0) && Board[i][y] == Board[x][y]; i-- {
			if BoardVisited[i][y] == 0 {
				fmt.Print("==Check first x-\n")
				nb, _ := check_three_free([]int{i, y}, player, sens)
				if nb > limit {
					nbr++
				}
				BoardVisited[i][y] = 1
			}
		}
	}
	if sens != 5 && sens != 6 {
		fmt.Print("=====Check les sens \\ \n")
		for x, y = coord[0], coord[1]; x >= 0 && y >= 0 && Board[x][y] == Board[xori][yori]; {
			if BoardVisited[x][y] == 0 {
				fmt.Print("==Check first x- y-\n")
				nb, _ := check_three_free([]int{x, y}, player, sens)
				if nb > limit {
					nbr++
				}
				BoardVisited[x][y] = 1
			}
			x--
			y--
		}

		for x, y = coord[0], coord[1]; x <= 19 && y <= 19 && Board[x][y] == Board[xori][yori]; {
			if BoardVisited[x][y] == 0 {
				fmt.Print("==Check first x+ y+\n")
				nb, _ := check_three_free([]int{x, y}, player, sens)
				if nb > limit {
					nbr++
				}
				BoardVisited[x][y] = 1
			}
			x++
			y++
		}
	}

	if sens != 7 && sens != 8 {
		fmt.Print("=====Check les sens / \n")
		for x, y = coord[0], coord[1]; x >= 0 && y <= 19 && Board[x][y] == Board[xori][yori]; {
			if BoardVisited[x][y] == 0 {
				fmt.Print("==Check first x- y+\n")
				nb, _ := check_three_free([]int{x, y}, player, sens)
				if nb > limit {
					nbr++
				}
				BoardVisited[x][y] = 1
			}
			x--
			y++
		}
		for x, y = coord[0], coord[1]; x <= 19 && y >= 0 && Board[x][y] == Board[xori][yori]; {
			if BoardVisited[x][y] == 0 {
				fmt.Print("==Check first x+ y-\n")
				nb, _ := check_three_free([]int{x, y}, player, sens)
				if nb > limit {
					nbr++
				}
				BoardVisited[x][y] = 1
			}
			x++
			y--
		}
	}
	fmt.Printf("\nTotal find %d\n", nbr)
	if nbr >= 1 {
		return true
	}
	return false
}

func dual_three(coord []int, player int) bool {

	Board[coord[0]][coord[1]] = player

	sens, nbr := check_three_free(coord, player, 0)
	if nbr > 0 {
		fmt.Printf("==========Creation d'un double trois on check si yen a dautre sauf dans le sens %d\n", -1)
		if check_dual_three(coord, player, sens) == true {
			Board[coord[0]][coord[1]] = NONE
			return true
		}
	}
	Board[coord[0]][coord[1]] = NONE
	return false
}

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

func referee(coord []string, ws *websocket.Conn) (bool, bool, int) {

	for pl, _ := range players {
		if (pl.ws == ws) && pl.player_color != Turn {
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
			send_capture(capture(coordint), ws)
			win, who := check_win(coordint)
			return true, win, who
		}
	}
	return false, false, -1
}
