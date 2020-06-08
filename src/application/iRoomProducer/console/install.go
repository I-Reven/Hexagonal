package console

import (
	"github.com/I-Reven/Hexagonal/src/application/iRoomProducer"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Install struct {
	Installer core.Installer
}

func (i *Install) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "Install core pkg",
		Long:  "Install core pkg dependency",
		Run: func(cmd *cobra.Command, args []string) {
			err := i.Installer.Install()

			if err != nil {
				color.HiRed(err.Error())
			} else {
				color.HiGreen("Install Done")
			}
		},
	}
}
