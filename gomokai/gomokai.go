package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	con, error := net.Dial("tcp", "localhost:1113")
	if error != nil {
		fmt.Printf("Host not found: %s\n", error)
		os.Exit(1)
	}

	in, error := con.Write([]byte("test"))
	if error != nil {
		fmt.Printf("Error sending data: %s, in: %d\n", error, in)
		os.Exit(2)
	}

	fmt.Println("Connection OK")

}
