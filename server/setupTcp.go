package server

import (
	"io"
	"log"
	"net"
	"strconv"

	"github.com/sebzz2k2/krompton/config"
	"github.com/sebzz2k2/krompton/core"
)

func respond(cmd *core.KromptonCmd, c io.ReadWriter) error {
	var b []byte
	switch cmd.Cmd {
	case "PING":
		b = core.Encode("PONG")
	}
	_, err := c.Write(b)
	return err
}

func readCommand(c net.Conn) (*core.KromptonCmd, error) {
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return nil, err
	}
	decoded, err := core.DecodeArrayStr(buf[:n])
	if err != nil {
		return nil, err
	}
	return &core.KromptonCmd{
		Cmd:  decoded[0],
		Args: decoded[1:],
	}, nil

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
			cmd, err := readCommand(conn)
			concurrentUser--
			if err == io.EOF {
				break
			}
			respond(cmd, conn)
		}

	}
}
