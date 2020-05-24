package console

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/server"
	"github.com/I-Reven/Hexagonal/src/infrastructures/cli"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/spf13/cobra"
)

var (
	ServeCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start to listen and serve core pkg",
		Long:  "Start to listen and serve hexagonal core package",
		Run: func(cmd *cobra.Command, args []string) {
			server.Listen()
		},
	}
)

//Cli Command line interface
func Cli() {
	err := cli.Execute(func(c *cobra.Command) {
		c.AddCommand(cli.VersionCmd)
		c.AddCommand(ServeCmd)
	})

	if err != nil {
		logger.Fatal(err)
	}
}
