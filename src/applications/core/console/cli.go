package console

import (
	"github.com/I-Reven/Hexagonal/src/applications/core"
	"github.com/I-Reven/Hexagonal/src/applications/core/job"
	"github.com/I-Reven/Hexagonal/src/applications/core/server"
	"github.com/I-Reven/Hexagonal/src/infrastructures/cli"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/fatih/color"
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

	InstallCmd = &cobra.Command{
		Use:   "install",
		Short: "Install core pkg",
		Long:  "Install core pkg dependency",
		Run: func(cmd *cobra.Command, args []string) {
			err := core.Install()

			if err != nil {
				color.HiRed(err.Error())
			} else {
				color.HiGreen("Install Done")
			}
		},
	}

	CronCmd = &cobra.Command{
		Use:   "cron",
		Short: "Run core pkg cron",
		Long:  "Run core pkg cron jobs",
		Run: func(cmd *cobra.Command, args []string) {
			job.Cron()
		},
	}
)

func Cli() {
	err := cli.Execute(func(c *cobra.Command) {
		c.AddCommand(cli.VersionCmd)
		c.AddCommand(ServeCmd)
		c.AddCommand(InstallCmd)
		c.AddCommand(CronCmd)
	})

	if err != nil {
		logger.Fatal(err)
	}
}
