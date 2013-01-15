package main

import (
	"log"
	"net"
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
			break
		}
		buff := string(data[0 : n-1])
		log.Printf("Receive '%s'", buff)
		parser(buff, con)
	}
}

func DialServ() (con net.Conn) {
	con, error := net.Dial("tcp", "localhost:1113")
	if error != nil {
		log.Printf("Host not found: %s\n", error)
	}
	log.Println("Connected to server")
	return con
}
