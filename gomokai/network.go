package main

import (
	"log"
	"net"
	"os"
	"strings"
)

func Send(buff string, con net.Conn) {
	in, error := con.Write([]byte(buff + "\n"))
	if error != nil {
		log.Printf("Error sending data: %s, in: %d\n", error, in)
	}
	log.Printf("Sending : '%s'", buff)
}

func HandleRead(con net.Conn) {
	var data = make([]byte, 70)

	for {
		n, err := con.Read(data)
		if err != nil {
			log.Println(err)
			os.Exit(11)
		}
		buff := string(data[0 : n-1])

		all_msg := strings.Split(buff, "\n")

		for _, msg := range all_msg {
			log.Printf("Receive '%s'", msg)
			parser(msg, con)
		}
	}
}

func DialServ(addr string) (con net.Conn) {
	con, error := net.Dial("tcp", addr)
	if error != nil {
		log.Printf("Host not found: %s\n", error)
		os.Exit(11)
	}
	log.Printf("Connected to server at %s", addr)
	return con
}
