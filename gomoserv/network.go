package main

import (
	"log"
	"strconv"
)

func send(buf string, c Connection) {
	log.Printf("send : '%s'", buf)
	if c.ws != nil {
		ws_send(buf, c.ws)
		return
	}
	c.s.Write([]byte(buf + "\n"))
}

func sendboard(s Connection) {

	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			buf := strconv.Itoa(x) + "," + strconv.Itoa(y)
			if Board[x][y] != NONE {
				buf = "ADD " + strconv.Itoa(x) + " " + strconv.Itoa(y)
				send(buf, s)
			}
		}
		log.Print("\n")
	}
}

func send_capture(pow [][]int, c Connection) {
	var buff string
	flag := 0

	if pow != nil {
		for key := range pow {
			if flag == 0 {
				buff = "REM "
			}
			if pow[key] != nil {
				buff += strconv.Itoa(pow[key][0]) + " " + strconv.Itoa(pow[key][1]) + " "
				if Board[pow[key][0]][pow[key][1]] == BLACK {
					WPOW += 1
				} else {
					BPOW += 1
				}
				Board[pow[key][0]][pow[key][1]] = NONE
				if flag == 1 {
					flag = 0
					for pl, _ := range players {
						send(buff, pl)
					}
					continue
				}
				flag++
			}

		}
	}
}
