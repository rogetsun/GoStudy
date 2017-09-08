package server

import (
	"GoStudy/uvsocket/server/analysis/byteAnalysis"
	"fmt"
	"net"
	"os"
)

func Listen(host string, port string) {
	fmt.Println("Initiating server... (Ctrl-C to stop)")

	remote := host + ":" + port
	lis, err := net.Listen("tcp", remote)
	defer lis.Close()

	if err != nil {
		fmt.Println("Error when listen: ", remote)
		os.Exit(-1)
	}
	fmt.Printf("server start to listenning %v:%v\n", host, port)
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("Error accepting client: ", err.Error())
			os.Exit(0)
		}
		fmt.Println("New connection: ", conn.RemoteAddr())
		go byteAnalysis.ByteRead(conn)
	}
}
