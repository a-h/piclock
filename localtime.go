package piclock

import "time"

// NewLocalTime creates a new local time source.
func NewLocalTime(initial time.Time) *LocalTime {
	lt := &LocalTime{
		t: initial,
		c: make(chan time.Duration),
	}
	go func() {
		previous := time.Now()
		tick := time.NewTicker(time.Millisecond * 50)
		for {
			select {
			case delta := <-lt.c:
				lt.t = lt.t.Add(delta)
			case current := <-tick.C:
				lt.t = lt.t.Add(current.Sub(previous))
				previous = current
			}
		}
	}()
	return lt
}

// LocalTime is the time local to the program.
type LocalTime struct {
	t time.Time
	c chan time.Duration
}

// Now returns the current time.
func (t *LocalTime) Now() time.Time {
	return t.t
}

// Add time to the clock.
func (t *LocalTime) Add(d time.Duration) {
	t.c <- d
}
