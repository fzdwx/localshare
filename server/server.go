package server

import (
	"fmt"
	"github.com/lxzan/gws"
	"localshare/web"
	"net/http"
	"time"
)

type Server struct {
	mux *http.ServeMux

	// dev mode
	// web assets will be reloaded every time the page is refreshed
	dev bool
	// port the server listens on
	port int
}

func Serve(opts ...Option) error {
	s := &Server{
		mux:  http.NewServeMux(),
		port: 9408,
		dev:  false,
	}

	for _, opt := range opts {
		opt.apply(s)
	}

	s.mount()
	return s.serve()
}

func (s *Server) serve() error {
	fmt.Printf("Server started at http://localhost:%d \n", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)
}

const (
	PingInterval = 5 * time.Second
	PingWait     = 10 * time.Second
)

func (s *Server) OnOpen(socket *gws.Conn) {
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait))
}

func (s *Server) OnClose(socket *gws.Conn, err error) {}

func (s *Server) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait))
	_ = socket.WritePong(nil)
}

func (s *Server) OnPong(socket *gws.Conn, payload []byte) {}

func (s *Server) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()
	socket.WriteMessage(message.Opcode, message.Bytes())
}

func (s *Server) mount() {
	var upgrader = gws.NewUpgrader(s, &gws.ServerOption{
		ParallelEnabled:   true,                                 // Parallel message processing
		Recovery:          gws.Recovery,                         // Exception recovery
		PermessageDeflate: gws.PermessageDeflate{Enabled: true}, // Enable compression
	})

	s.mux.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := upgrader.Upgrade(writer, request)
		if err != nil {
			return
		}
		go func() {
			socket.ReadLoop() // Blocking prevents the context from being GC.
		}()
	})

	if s.dev {
		s.mux.Handle("/", http.FileServer(http.Dir("web")))
		fmt.Printf("Dev mode enabled\n")
	} else {
		s.mux.Handle("/", http.FileServerFS(web.Dist))
	}
}
