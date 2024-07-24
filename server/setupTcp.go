package server

import (
	"bufio"
	"github.com/sebzz2k2/krompton/config"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

func respond(command string, conn net.Conn) error {
	// writes to tcp connection
	_, err := conn.Write([]byte(command))
	if err != nil {
		return err
	}
	return nil
}

func tokenize(input string) []string {
	input = strings.TrimSpace(input)
	tokens := strings.Fields(input)
	return tokens
}

func handleInput(reader *bufio.Reader) ([]string, error) {
	// reads from tcp connection
	message, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	tokens := tokenize(message)
	return tokens, nil
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
			panic(err)
		}
		concurrentUser++
		log.Printf("Accepted connection from %s", conn.RemoteAddr())

		for {
			reader := bufio.NewReader(conn)
			cmds, err := handleInput(reader)
			if err != nil {
				conn.Close()
				concurrentUser--
				log.Printf("Connection closed by %s", conn.RemoteAddr())
				if err == io.EOF {
					break
				}
			}

			if cmds[0] == "ping" {
				respond("pong\n", conn)
			} else {
				respond("Unknown Command\n", conn)
			}
		}
	}
}
