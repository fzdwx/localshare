package server

type Option interface {
	apply(*Server)
}

type WithDev bool

type WithPort int

func (w WithDev) apply(server *Server) {
	server.dev = bool(w)
}
func (w WithPort) apply(server *Server) {
	server.port = int(w)
}
