package job

import "math/rand"

func Cron() {
	EveryMinute()

	if rand.Intn(5) == 1 {
		EveryFiveMinute()
	}

	if rand.Intn(15) == 1 {
		EveryFifineMinute()
	}

	if rand.Intn(30) == 1 {
		EveryThirtyMinute()
	}

	if rand.Intn(60) == 1 {
		EveryHour()
	}

	if rand.Intn(360) == 1 {
		EverySixHour()
	}

	if rand.Intn(720) == 1 {
		EveryTwelveHour()
	}

	if rand.Intn(1440) == 1 {
		EveryDay()
	}

	if rand.Intn(10080) == 1 {
		EveryWeek()
	}

	if rand.Intn(43200) == 1 {
		EveryMonth()
	}
}

func EveryMinute() {

}

func EveryFiveMinute() {

}

func EveryFifineMinute() {

}

func EveryThirtyMinute() {

}

func EveryHour() {

}

func EverySixHour() {

}

func EveryTwelveHour() {

}

func EveryDay() {

}

func EveryWeek() {

}

func EveryMonth() {

}
