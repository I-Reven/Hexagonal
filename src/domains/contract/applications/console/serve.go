package console

import "github.com/spf13/cobra"

type Serve interface {
	Command() *cobra.Command
}
