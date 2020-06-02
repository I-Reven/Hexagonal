package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func (*CLI) VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hexagonal v" + os.Getenv("APP_VERSION"))
		},
	}
}
