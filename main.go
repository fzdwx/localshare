package main

import "localshare/server"

func main() {
	if err := server.Serve(
		server.WithPort(9999),
		server.WithDev(),
	); err != nil {
		panic(err)
	}
}
