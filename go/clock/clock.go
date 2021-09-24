package clock

import "fmt"

type Clock struct {
	Hour   int
	Minute int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.Hour, c.Minute+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return New(c.Hour, c.Minute-minutes)
}

func New(hour, minute int) Clock {
	hour, minute = norm(hour, minute, 60)
	_, hour = norm(0, hour, 24)
	return Clock{
		Hour:   hour,
		Minute: minute,
	}
}

func norm(hi, lo, base int) (nhi, nlo int) {
	if lo < 0 {
		n := (-lo-1)/base + 1
		hi -= n
		lo += n * base
	}
	if lo >= base {
		n := lo / base
		hi += n
		lo -= n * base
	}
	return hi, lo
}
