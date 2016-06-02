package clock

import "fmt"

const testVersion = 4
const minutesInDay = 60 * 24

type Clock struct {
	mins int
}

func convertMins(mins int) int {
	estimated := mins % minutesInDay
	if estimated < 0 {
		estimated = minutesInDay + estimated
	}

	return estimated
}

func New(hour, minute int) Clock {
	return Clock{convertMins(hour*60 + minute)}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.mins/60, c.mins%60)
}

func (c Clock) Add(minutes int) Clock {
	c.mins = convertMins(minutes + c.mins)

	return c
}
