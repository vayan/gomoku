package main

import (
	"log"
)

func main() {
	log.Print("HI")
	con := DialServ()
	go HandleRead(con)
	Send("CONNECT CLIENT", con)
	select {}
}
