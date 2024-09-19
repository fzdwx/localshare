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

func WithDev(dev bool) Option {
	return newOption(func(s *Server) {
		s.dev = dev
	})
}

func WithPort(port int) Option {
	return newOption(func(s *Server) {
		s.port = port
	})
}
