package alarmclock

import (
	"fmt"
	"time"
)

// Screen in the display.
type Screen interface {
	Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string)
}

// NewState inititalises the state.
func NewState() State {
	return State{
		Time: time.Now,
	}
}

// State of the display.
type State struct {
	Time            func() time.Time
	AlarmEnabled    bool
	AlarmTimeHour   int
	AlarmTimeMinute int
	Volume          int
}

type home struct{}

func (screen *home) Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.
	if up {
		current = clockScreen
		return
	}

	// Update the screen.
	line1 = "Eddie's alarm"
	line2 = "clock"
	return
}

type clock struct{}

func (screen *clock) Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.
	if down {
		current = HomeScreen
		return
	}
	if up {
		current = calendarScreen
		return
	}

	// Update the screen.
	line1 = "Time:"
	line2 = fmt.Sprintf("%s ", s.Time().Format("15:04:05"))
	return
}

type calendar struct{}

func (screen *calendar) Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.
	if down {
		current = clockScreen
		return
	}

	// Update the screen.
	line1 = "Date:"
	line2 = fmt.Sprintf("%s ", s.Time().Format("Mon 02/01/2006"))
	return
}

// Screens of the system.

// HomeScreen of the system.
var HomeScreen Screen = &home{}
var clockScreen Screen = &clock{}
var calendarScreen Screen = &calendar{}
