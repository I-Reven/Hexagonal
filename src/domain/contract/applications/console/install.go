package console

import "github.com/spf13/cobra"

type Install interface {
	Command() *cobra.Command
}
