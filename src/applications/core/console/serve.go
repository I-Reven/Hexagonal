package console

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/listener"
	"github.com/spf13/cobra"
)

type Serve struct {
	Listener listener.Listener
}

func (c Serve) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start to listen and serve core pkg",
		Long:  "Start to listen and serve hexagonal core package",
		Run: func(cmd *cobra.Command, args []string) {
			c.Listener.Listen()
		},
	}
}
