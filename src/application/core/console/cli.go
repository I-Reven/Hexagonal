package console

import (
	"github.com/I-Reven/Hexagonal/src/application/core"
	"github.com/I-Reven/Hexagonal/src/framework/cli"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	_ "github.com/mattn/go-colorable"
	"github.com/spf13/cobra"
)

type Cli struct {
	Log     logger.Log
	Core    core.Kernel
	CLI     cli.CLI
	Serve   Serve
	Install Install
	Cron    Cron
}

func (c *Cli) Boot() {
	c.Core.Boot()
	c.Cli()
}

func (c *Cli) Cli() {
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
