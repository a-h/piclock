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
		current = alarmStateScreen
		return
	}

	// Update the screen.
	line1 = "Eddie's alarm"
	line2 = "clock"
	return
}

type alarmState struct{}

func (screen *alarmState) Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.
	if down {
		current = HomeScreen
		return
	}

	// Update the screen.
	line1 = "Alarm state"
	line2 = fmt.Sprintf("%s %v", s.Time().Format("15:04:05"), s.AlarmEnabled)
	return
}

// Screens of the system.

// HomeScreen of the system.
var HomeScreen Screen = &home{}
var alarmStateScreen Screen = &alarmState{}
