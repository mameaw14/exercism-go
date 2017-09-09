package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	time := hour * 60 + minute
	time %= 60 * 24
	if time < 0 {
		time += 60 * 24
	}
	hour = (time / 60) % 24
	minute = time % 60

	return Clock{hour, minute}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
	clock.minute += minutes
	return New(clock.hour,clock.minute)
}
