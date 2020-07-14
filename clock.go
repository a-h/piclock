package piclock

import (
	"fmt"
	"strconv"
	"time"
)

// Screen in the display.
type Screen interface {
	Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string)
}

// NewState inititalises the state.
func NewState() State {
	return State{
		Time: NewLocalTime(time.Now()),
	}
}

// State of the display.
type State struct {
	Time            *LocalTime
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
	line2 = fmt.Sprintf("%s ", s.Time.Now().Format("15:04:05"))
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
	if up {
		current = alarmStateScreen
		return
	}

	// Update the screen.
	line1 = "Date:"
	line2 = fmt.Sprintf("%s ", s.Time.Now().Format("Mon 02/01/2006"))
	return
}

type alarmState struct{}

func onOff(v bool) string {
	if v {
		return "On"
	}
	return "Off"
}

func (screen *alarmState) Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.
	if down {
		current = calendarScreen
		return
	}

	if up {
		current = settingsScreen
		return
	}
	// Update the screen.
	line1 = "Alarm state:"
	line2 = fmt.Sprintf("%02d:%02d %v", s.AlarmTimeHour, s.AlarmTimeMinute, onOff(s.AlarmEnabled))
	return
}

type settings struct{}

func (screen *settings) Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.
	if down {
		current = alarmStateScreen
		return
	}
	if right {
		current = settingsSetTimeScreen
		return
	}
	// Update the screen.
	line1 = "Settings:"
	line2 = "open~"
	return
}

type settingsSetTime struct{}

func (screen *settingsSetTime) Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.
	if right {
		current = settingsSettingTimeHourScreen
		return
	}
	if left {
		current = settingsScreen
		return
	}
	if up {
		current = settingsSetAlarmTimeScreen
		return
	}
	// Update the screen.
	line1 = "Settings:"
	line2 = "Set Time~"
	return
}

type settingsSetAlarmTime struct{}

func (screen *settingsSetAlarmTime) Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.

	if down {
		current = settingsSetTimeScreen
		return
	}
	if right {
		current = settingsSettingAlarmTimeHourScreen
		return
	}
	// Update the screen.
	line1 = "Settings:"
	line2 = "Set Alarm Time~"
	return
}

type settingsSettingAlarmTimeHour struct{}

func (screen *settingsSettingAlarmTimeHour) Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.

	if left {
		current = settingsSetTimeScreen
		return
	}
	if right {
		current = settingsSettingAlarmTimeMinScreen
		return
	}
	if up {
		if s.AlarmTimeHour >= 24 {
			s.AlarmTimeHour = 0
		} else {
			s.AlarmTimeHour++
		}
		return
	}
	if down {
		if s.AlarmTimeHour <= 0 {
			s.AlarmTimeHour = 24
		} else {
			s.AlarmTimeHour--
		}
	}
	// Update the screen.
	line1 = "Set Alarm Hour:"
	line2 = strconv.Itoa(s.AlarmTimeHour) + ":" + strconv.Itoa(s.AlarmTimeMinute)
	return
}

type settingsSettingAlarmTimeMin struct{}

func (screen *settingsSettingAlarmTimeMin) Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.

	if left {
		current = settingsSettingAlarmTimeHourScreen
		return
	}
	if up {
		if s.AlarmTimeMinute >= 60 {
			s.AlarmTimeMinute = 0
		} else {
			s.AlarmTimeHour++
		}
		return
	}
	if down {
		if s.AlarmTimeHour <= 0 {
			s.AlarmTimeHour = 60
		} else {
			s.AlarmTimeHour--
		}
	}
	// Update the screen.
	line1 = "Set Alarm Hour:"
	line2 = strconv.Itoa(s.AlarmTimeHour) + ":" + strconv.Itoa(s.AlarmTimeMinute)
	return
}

type settingsSettingTimeHour struct{}

func (screen *settingsSettingTimeHour) Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.
	if down {
		switch state.Time.Now().Hour() {
		case 0:
			state.Time.Add(23 * time.Hour)
		default:
			state.Time.Add(-time.Hour)
		}
	}
	if up {
		switch state.Time.Now().Hour() {
		case 23:
			state.Time.Add(-23 * time.Hour)
		default:
			state.Time.Add(time.Hour)
		}
	}
	if right {
		current = settingsSettingTimeMinScreen
		return
	}

	if left {
		current = settingsSetTimeScreen
		return
	}
	// Update the screen.
	line1 = "Set hour:"
	line2 = fmt.Sprintf("%s ", s.Time.Now().Format("15:04:05"))
	return
}

type settingsSettingTimeMin struct{}

func (screen *settingsSettingTimeMin) Update(s State, up, down, left, right bool) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.
	if down {
		switch state.Time.Now().Minute() {
		case 0:
			state.Time.Add(59 * time.Minute)
		default:
			state.Time.Add(-time.Minute)
		}
	}
	if up {
		switch state.Time.Now().Minute() {
		case 59:
			state.Time.Add(-59 * time.Minute)
		default:
			state.Time.Add(time.Minute)
		}
	}

	if left {
		current = settingsSettingTimeHourScreen
		return
	}
	// Update the screen.
	line1 = "Set minute:"
	line2 = fmt.Sprintf("%s ", s.Time.Now().Format("15:04:05"))
	return
}

// Screens of the system.

// HomeScreen is the first screen that gets loaded.
var HomeScreen Screen = &home{}
var clockScreen Screen = &clock{}
var calendarScreen Screen = &calendar{}
var alarmStateScreen = &alarmState{}
var settingsScreen = &settings{}
var settingsSetTimeScreen = &settingsSetTime{}
var settingsSettingTimeHourScreen = &settingsSettingTimeHour{}
var settingsSettingTimeMinScreen = &settingsSettingTimeMin{}
var settingsSetAlarmTimeScreen = &settingsSetAlarmTime{}
var settingsSettingAlarmTimeHourScreen = &settingsSettingAlarmTimeHour{}
var settingsSettingAlarmTimeMinScreen = &settingsSettingAlarmTimeMin{}
