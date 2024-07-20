package server

import (
	"github.com/sebzz2k2/krompton/config"
	"io"
	"log"
	"net"
	"strconv"
)

func respond(command string, conn net.Conn) error {
	// writes to tcp connection
	_, err := conn.Write([]byte(command))
	if err != nil {
		return err
	}
	return nil
}

func readCommand(conn net.Conn) (string, error) {
	// reads from tcp connection
	var buffer []byte = make([]byte, 512)
	n, err := conn.Read(buffer[:])
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}

func RunSyncTcp() {
	var concurrentUser = 0
	listener, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			if err != nil {
				panic(err)
			}
		}
		concurrentUser++
		log.Printf("Accepted connection from %s", conn.RemoteAddr())

		for {
			cmd, err := readCommand(conn)
			if err != nil {
				conn.Close()
				concurrentUser--
				log.Printf("Connection closed by %s", conn.RemoteAddr())
				if err == io.EOF {
					break
				}
			}
			log.Printf("Command: %s", cmd)
			if err = respond(cmd, conn); err != nil {
				log.Println(err)
			}
		}
	}
}
