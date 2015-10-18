package clock

import "fmt"

//TestVersion is the number of testing version
const (
	TestVersion    = 2
	minutesPerDay  = 1440
	minutesPerHour = 60
)

//Clock is a simple int
type Clock int

//positiveClock checks if the minutes in the clock is negative. If so, it normalizes it.
func positiveClock(c Clock) Clock {
	if c < 0 {
		c += minutesPerDay
	}
	return c
}

//Time creates a Clock and converts the hour in minutes
func Time(hour, minute int) Clock {
	c := Clock((hour*minutesPerHour + minute) % minutesPerDay)
	return positiveClock(c)
}

//String() is a stringer method. It returns the string version of the Clock in the form 08:00
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/minutesPerHour, c%minutesPerHour)
}

//Add add the given minutes to the Clock
func (c Clock) Add(minutes int) Clock {
	c = (c + Clock(minutes)) % minutesPerDay
	return positiveClock(c)
}
