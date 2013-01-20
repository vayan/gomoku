package main

import ()

func check_dual_three(coord Coord, player int, sens int, board [][]int) bool {
	x := coord.x
	y := coord.y

	xori := coord.x
	yori := coord.y

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

		for i := y; (i >= 0) && board[x][i] == board[x][y]; i-- {
			if BoardVisited[x][i] == 0 {

				nb, _ := check_three_free(Coord{x, i}, player, sens, board)
				if nb > limit {
					nbr++
				}
				BoardVisited[x][i] = 1
			}
		}
		for i := y; (i <= 19) && board[x][i] == board[x][y]; i++ {
			if BoardVisited[x][i] == 0 {

				nb, _ := check_three_free(Coord{x, i}, player, sens, board)
				if nb > limit {
					nbr++
				}
				BoardVisited[x][i] = 1
			}
		}
	}
	if sens != 3 && sens != 4 {

		for i := x; (i <= 19) && board[i][y] == board[x][y]; i++ {
			if BoardVisited[i][y] == 0 {

				nb, _ := check_three_free(Coord{i, y}, player, sens, board)
				if nb > limit {
					nbr++
				}
				BoardVisited[i][y] = 1
			}
		}
		for i := x; (i >= 0) && board[i][y] == board[x][y]; i-- {
			if BoardVisited[i][y] == 0 {

				nb, _ := check_three_free(Coord{i, y}, player, sens, board)
				if nb > limit {
					nbr++
				}
				BoardVisited[i][y] = 1
			}
		}
	}
	if sens != 5 && sens != 6 {

		for x, y = coord.x, coord.y; x >= 0 && y >= 0 && board[x][y] == board[xori][yori]; {
			if BoardVisited[x][y] == 0 {

				nb, _ := check_three_free(Coord{x, y}, player, sens, board)
				if nb > limit {
					nbr++
				}
				BoardVisited[x][y] = 1
			}
			x--
			y--
		}

		for x, y = coord.x, coord.y; x <= 19 && y <= 19 && board[x][y] == board[xori][yori]; {
			if BoardVisited[x][y] == 0 {

				nb, _ := check_three_free(Coord{x, y}, player, sens, board)
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

		for x, y = coord.x, coord.y; x >= 0 && y <= 19 && board[x][y] == board[xori][yori]; {
			if BoardVisited[x][y] == 0 {

				nb, _ := check_three_free(Coord{x, y}, player, sens, board)
				if nb > limit {
					nbr++
				}
				BoardVisited[x][y] = 1
			}
			x--
			y++
		}
		for x, y = coord.x, coord.y; x <= 19 && y >= 0 && board[x][y] == board[xori][yori]; {
			if BoardVisited[x][y] == 0 {

				nb, _ := check_three_free(Coord{x, y}, player, sens, board)
				if nb > limit {
					nbr++
				}
				BoardVisited[x][y] = 1
			}
			x++
			y--
		}
	}
	if nbr >= 1 {
		return true
	}
	return false
}

func check_three_free(coord Coord, player int, sens int, board [][]int) (int, int) {
	x := coord.x
	y := coord.y

	xori := coord.x
	yori := coord.y

	nbr := 0
	nb_free_hori := 0
	nb_free_verti := 0
	nb_free_slash := 0
	nb_free_anti := 0
	tmp := nbr
	limit := 4

	if sens != 1 && sens != 2 {
		for i, v, stone := y, 0, 0; (i >= 0) && (board[x][i] == player || board[x][i] == NONE) && v <= 3; i-- {
			if board[x][i] == player {
				stone++
				nb_free_hori++
			}
			if stone == limit && i-1 >= 0 && yori+1 <= 19 && board[x][yori+1] == NONE && board[x][i-1] == NONE {

				nbr++
				return 1, 1
			}
			v++
		}
		for i, v, stone := y, 0, 0; (i <= 19) && (board[x][i] == player || board[x][i] == NONE) && v <= 3; i++ {
			if board[x][i] == player {
				stone++
				nb_free_hori++
			}
			if stone == limit && i+1 <= 19 && yori-1 >= 0 && board[x][yori-1] == NONE && board[x][i+1] == NONE {

				nbr++
				return 2, 1
			}
			v++
		}
		if nb_free_hori == limit && tmp == nbr {

			return 2, 1
		}
	}

	if sens != 3 && sens != 4 {
		tmp = nbr
		for i, v, stone := x, 0, 0; (i <= 19) && (board[i][y] == player || board[i][y] == NONE) && v <= 3; i++ {
			if board[i][y] == player {
				stone++
				nb_free_verti++
			}
			if stone == limit && i+1 <= 19 && xori-1 >= 0 && board[xori-1][y] == NONE && board[i+1][y] == NONE {

				nbr++
				return 3, 1
			}
			v++

		}
		for i, v, stone := x, 0, 0; (i >= 0) && (board[i][y] == player || board[i][y] == NONE) && v <= 3; i-- {
			if board[i][y] == player {
				stone++
				nb_free_verti++
			}
			if stone == limit && i-1 >= 0 && xori+1 <= 19 && board[xori+1][y] == NONE && board[i-1][y] == NONE {

				nbr++
				return 4, 1
			}
			v++
		}
		if nb_free_verti == limit && tmp == nbr {

			return 4, 1
		}
	}
	if sens != 5 && sens != 6 {
		tmp = nbr
		for x, y, v, stone := coord.x, coord.y, 0, 0; x >= 0 && y >= 0 && (board[x][y] == player || board[x][y] == NONE) && v <= 3; {
			if board[x][y] == player {
				stone++
				nb_free_slash++
			}
			if stone == limit && x-1 >= 0 && y-1 >= 0 && xori+1 <= 19 && yori+1 <= 19 && board[xori+1][yori+1] == NONE && board[x-1][y-1] == NONE {

				nbr++
				return 5, 1
			}
			v++
			x--
			y--
		}
		for x, y, v, stone := coord.x, coord.y, 0, 0; x <= 19 && y <= 19 && (board[x][y] == player || board[x][y] == NONE) && v <= 3; {
			if board[x][y] == player {
				stone++
				nb_free_slash++
			}
			if stone == limit && x+1 <= 19 && y+1 <= 19 && xori-1 >= 0 && yori-1 >= 0 && board[xori-1][yori-1] == NONE && board[x+1][y+1] == NONE {

				nbr++
				return 6, 1
			}
			v++
			x++
			y++
		}
		if nb_free_slash == limit && tmp == nbr {

			return 6, 1
		}
	}
	if sens != 7 && sens != 8 {
		tmp = nbr
		for x, y, v, stone := coord.x, coord.y, 0, 0; x >= 0 && y <= 19 && (board[x][y] == player || board[x][y] == NONE) && v <= 3; {
			if board[x][y] == player {
				stone++
				nb_free_anti++
			}
			if stone == limit && x-1 >= 0 && y+1 <= 19 && xori+1 <= 19 && yori-1 >= 0 && board[xori+1][yori-1] == NONE && board[x-1][y+1] == NONE {

				nbr++
				return 7, 1
			}
			v++
			x--
			y++
		}

		for x, y, v, stone := coord.x, coord.y, 0, 0; x <= 19 && y >= 0 && (board[x][y] == player || board[x][y] == NONE) && v <= 3; {
			if board[x][y] == player {
				stone++
				nb_free_anti++
			}
			if stone == limit && x+1 <= 19 && y-1 >= 0 && xori-1 >= 0 && yori+1 <= 19 && board[xori-1][yori+1] == NONE && board[x+1][y-1] == NONE {

				nbr++
				return 8, 1
			}
			v++
			x++
			y--
		}
		if nb_free_anti == limit && tmp == nbr {

			return 8, 1
		}
	}
	return 0, 0
}

func dual_three(coord Coord, player int, board [][]int) bool {

	if DOUBLE_3 == 0 {
		return false
	}
	board[coord.x][coord.y] = player

	sens, nbr := check_three_free(coord, player, 0, board)
	if nbr > 0 {
		if check_dual_three(coord, player, sens, board) == true {
			board[coord.x][coord.y] = NONE
			return true
		}
	}
	board[coord.x][coord.y] = NONE
	return false
}
