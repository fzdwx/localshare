package main

import (
	"github.com/spf13/cobra"
	"localshare/server"
	"os"
)

func main() {
	var port int
	var cmd = &cobra.Command{
		Use:  "localshare",
		Long: "Localshare is a simple file sharing server",
		RunE: func(cmd *cobra.Command, args []string) error {
			dev := os.Getenv("LOCALSHARE_DEV")
			return server.Serve(
				server.WithPort(9999),
				server.WithDev(dev == "true"),
			)
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 9999, "Port to listen on")

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
