package cli

import "github.com/spf13/cobra"

type Cli interface {
	Execute(addCommand func(*cobra.Command)) error
	VersionCmd() *cobra.Command
}
