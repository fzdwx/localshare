package server

import (
	"fmt"
	"github.com/lxzan/gws"
	"localshare/web"
	"net/http"
	"sync"
)

type Server struct {
	mux      *http.ServeMux
	users    map[string]*gws.Conn
	usersMux sync.RWMutex

	// dev mode
	// web assets will be reloaded every time the page is refreshed
	dev bool
	// port the server listens on
	port int
}

func Serve(opts ...Option) error {
	s := &Server{
		mux:      http.NewServeMux(),
		users:    make(map[string]*gws.Conn),
		usersMux: sync.RWMutex{},
		port:     9408,
		dev:      false,
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

func (s *Server) AddUser(userID string, socket *gws.Conn) {
	s.usersMux.Lock()
	defer s.usersMux.Unlock()
	s.users[userID] = socket
}

func (s *Server) RemoveUser(userID string) {
	s.usersMux.Lock()
	defer s.usersMux.Unlock()
	delete(s.users, userID)
}

func (s *Server) Broadcast(excludeMsgUser bool, msg *Message) {
	s.usersMux.RLock()
	defer s.usersMux.RUnlock()

	for userID, conn := range s.users {
		if excludeMsgUser && userID == msg.UserID {
			continue
		}
		conn.WriteMessage(gws.OpcodeText, msg.raw)
	}
}
