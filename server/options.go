package server

type Option interface {
	apply(*Server)
}

type optionFunc func(*Server)

func (f optionFunc) apply(s *Server) {
	f(s)
}

func newOption(f func(*Server)) Option {
	return optionFunc(f)
}

func WithDev() Option {
	return newOption(func(s *Server) {
		s.dev = true
	})
}

func WithPort(port int) Option {
	return newOption(func(s *Server) {
		s.port = port
	})
}
