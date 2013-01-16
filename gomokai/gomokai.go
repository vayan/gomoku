package main

import (
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func Atoi(st string) int {
	ret, _ := strconv.Atoi(st)
	return ret
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func get_opos_turn(turn int) int {
	if turn == BLACK {
		return WHITE
	}
	return BLACK
}

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func main() {
	runtime.GOMAXPROCS(MaxParallelism())
	rand.Seed(time.Now().UTC().UnixNano())
	log.Print("HI")
	con := DialServ()
	go HandleRead(con)
	Send("CONNECT CLIENT", con)
	select {}
}
