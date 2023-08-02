package lib

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type TCPServer struct {
	host     string
	port     int
	handlers map[string]func(string) string
}

func NewTCPServer(host string, port int) *TCPServer {
	return &TCPServer{
		host:     host,
		port:     port,
		handlers: make(map[string]func(string) string),
	}
}

func (t *TCPServer) RegisterHandler(name string, handler func(string) string) {
	t.handlers[strings.ToLower(name)] = handler
}

func (t *TCPServer) ListenAndServe() error {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", t.host, t.port))
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go t.handleConnection(conn)
	}
}

func (t *TCPServer) handleConnection(conn net.Conn) {
	fmt.Printf("[TCP][%s] OPEN\n", conn.RemoteAddr().String())
	defer func() {
		fmt.Printf("[TCP][%s] CLOSE\n", conn.RemoteAddr().String())
		conn.Close()
	}()

	s := bufio.NewScanner(conn)
	for s.Scan() {
		req := s.Text()
		cmd := strings.ToLower(strings.Split(req, " ")[0])

		if handler, ok := t.handlers[cmd]; ok {
			res := handler(req)
			fmt.Printf("[TCP][%s] %s -> %s\n", conn.RemoteAddr().String(), req, res)
			conn.Write([]byte(res))
		} else {
			fmt.Printf("[TCP][%s] Unknown command: %s\n", conn.RemoteAddr().String(), cmd)
		}
	}
}
