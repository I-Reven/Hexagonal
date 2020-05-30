package console

import (
	"github.com/I-Reven/Hexagonal/src/applications/core"
	"github.com/I-Reven/Hexagonal/src/infrastructures/cli"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	_ "github.com/mattn/go-colorable"
	"github.com/spf13/cobra"
)

type CLI struct {
	Log     logger.Log
	Core    core.Core
	CLI     cli.CLI
	Serve   Serve
	Install Install
	Cron    Cron
}

func (c CLI) Boot() {
	c.Core.Boot()
	c.Cli()
}

func (c CLI) Cli() {
	err := c.CLI.Execute(func(cobra *cobra.Command) {
		cobra.AddCommand(c.CLI.VersionCmd())
		cobra.AddCommand(c.Serve.Command())
		cobra.AddCommand(c.Install.Command())
		cobra.AddCommand(c.Cron.CronCmd())
	})

	if err != nil {
		c.Log.Fatal(err)
	}
}
