package console

import "github.com/spf13/cobra"

type Cron interface {
	CronCmd() *cobra.Command
}
