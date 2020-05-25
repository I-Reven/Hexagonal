package console

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/server"
	"github.com/I-Reven/Hexagonal/src/infrastructures/cli"
	"github.com/I-Reven/Hexagonal/src/infrastructures/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/cassandra/log"
	"github.com/fatih/color"
	"github.com/juju/errors"
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

	MigrateCmd = &cobra.Command{
		Use:   "migration",
		Short: "Migrate core pkg database",
		Long:  "Migrate core pkg cassandra database",
		Run: func(cmd *cobra.Command, args []string) {
			err := log.Log().Migrate()

			if err != nil {
				err = errors.NewNotSupported(err, "Can not migrate cassandra logs")
				color.HiRed(err.Error())
				logger.Panic(err)
			}

			color.HiGreen("Migration Done")
		},
	}
)

//Cli Command line interface
func Cli() {
	err := cli.Execute(func(c *cobra.Command) {
		c.AddCommand(cli.VersionCmd)
		c.AddCommand(ServeCmd)
		c.AddCommand(MigrateCmd)
	})

	if err != nil {
		logger.Fatal(err)
	}
}
