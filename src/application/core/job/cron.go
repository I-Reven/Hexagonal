package job

import "math/rand"

type CronJob struct{}

func (c CronJob) Cron() {
	c.everyMinute()

	if rand.Intn(5) == 1 {
		c.everyFiveMinute()
	}
}

func (c CronJob) everyMinute() {

}

func (c CronJob) everyFiveMinute() {

}
