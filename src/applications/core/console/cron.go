package console

import (
	"github.com/I-Reven/Hexagonal/src/applications/core/job"
	"github.com/spf13/cobra"
)

type Cron struct {
	CronJob job.CronJob
}

func (c *Cron) CronCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "cron",
		Short: "Run core pkg cron",
		Long:  "Run core pkg cron jobs",
		Run: func(cmd *cobra.Command, args []string) {
			c.CronJob.Cron()
		},
	}
}
