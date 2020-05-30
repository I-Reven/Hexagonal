package job

import "math/rand"

type CronJob struct{}

func (c CronJob) Cron() {
	c.everyMinute()

	if rand.Intn(5) == 1 {
		c.everyFiveMinute()
	}

	if rand.Intn(15) == 1 {
		c.everyFifineMinute()
	}

	if rand.Intn(30) == 1 {
		c.everyThirtyMinute()
	}

	if rand.Intn(60) == 1 {
		c.everyHour()
	}
}

func (c CronJob) everyMinute() {

}

func (c CronJob) everyFiveMinute() {

}

func (c CronJob) everyFifineMinute() {

}

func (c CronJob) everyThirtyMinute() {

}

func (c CronJob) everyHour() {

}
