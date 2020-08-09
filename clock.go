package piclock

import (
	"fmt"
	"strconv"
	"time"
)

// Screen in the display.
type Screen interface {
	Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string)
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
	AlarmType       int
}

type home struct{}

func (screen *home) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
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

func (screen *clock) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
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

func (screen *calendar) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
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
func alarmtypedecoder(v int) string {
	if v == 0 {
		return "SDCARD~"
	} else if v == 1 {
		return "SDCARD SHUFFLE"
	} else if v == 2 {
		return "GLADOS"
	}
	return "ERR"
}

func (screen *alarmState) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
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

func (screen *settings) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
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

func (screen *settingsSetTime) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
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

func (screen *settingsSetAlarmTime) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.

	if down {
		current = settingsSetTimeScreen
		return
	}
	if left {
		current = settingsScreen
		return
	}
	if right {
		current = settingsSettingAlarmTimeHourScreen
		return
	}
	if up {
		current = settingsSetAlarmToggleScreen
	}
	// Update the screen.
	line1 = "Settings:"
	line2 = "Set Alarm Time~"
	return
}

type settingsSettingAlarmTimeHour struct{}

func (screen *settingsSettingAlarmTimeHour) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.

	if left {
		current = settingsSetAlarmTimeScreen
		return
	}
	if right {
		current = settingsSettingAlarmTimeMinScreen
		return
	}
	if up {
		if state.AlarmTimeHour > 23 {
			state.AlarmTimeHour = 0
		} else {
			state.AlarmTimeHour++
		}
		return
	}
	if down {
		if state.AlarmTimeHour <= 0 {
			state.AlarmTimeHour = 23
		} else {
			state.AlarmTimeHour--
		}
		return
	}
	// Update the screen.
	line1 = "Set Alarm Hour:"
	if state.AlarmTimeMinute < 10 {
		line2 = strconv.Itoa(state.AlarmTimeHour) + ":" + "0" + strconv.Itoa(state.AlarmTimeMinute)
	} else {
		line2 = strconv.Itoa(state.AlarmTimeHour) + ":" + strconv.Itoa(state.AlarmTimeMinute)
	}
	return
}

type settingsSettingAlarmTimeMin struct{}

func (screen *settingsSettingAlarmTimeMin) Update(s State, up bool, down bool, left bool, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
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
		if state.AlarmTimeMinute > 59 {
			state.AlarmTimeMinute = 0
		} else {
			state.AlarmTimeMinute++
		}
		return
	}
	if down {
		if state.AlarmTimeMinute <= 0 {
			state.AlarmTimeMinute = 59
		} else {
			state.AlarmTimeMinute--
		}
		return
	}
	// Update the screen.
	line1 = "Set Alarm Minute:"
	if s.AlarmTimeMinute < 10 {
		line2 = strconv.Itoa(s.AlarmTimeHour) + ":" + "0" + strconv.Itoa(s.AlarmTimeMinute)
	} else {
		line2 = strconv.Itoa(s.AlarmTimeHour) + ":" + strconv.Itoa(s.AlarmTimeMinute)
	}
	return
}

type settingsSettingTimeHour struct{}

func (screen *settingsSettingTimeHour) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
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

func (screen *settingsSettingTimeMin) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
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

type settingsSetAlarmToggle struct{}

func (screen *settingsSetAlarmToggle) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.

	if down {
		current = settingsSetAlarmTimeScreen
		return
	}
	if left {
		current = settingsScreen
		return
	}
	if right {
		current = settingsSettingAlarmToggleScreen
		return
	}
	if up {
		current = settingsSetAlarmTypeScreen
		return
	}
	// Update the screen.
	line1 = "Settings:"
	line2 = "Toggle Alarm~"
	return
}

type settingsSettingAlarmToggle struct{}

func (screen *settingsSettingAlarmToggle) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.

	if down {
		state.AlarmEnabled = false
		return
	}
	if left {
		current = settingsScreen
		return
	}
	if up {
		state.AlarmEnabled = true
		return
	}
	// Update the screen.
	line1 = "Alarm:"
	line2 = onOff(state.AlarmEnabled)
	return
}

type settingsSetAlarmToggleType struct{}

func (screen *settingsSetAlarmToggleType) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.

	if down {
		current = settingsSetAlarmToggleScreen
		return
	}
	if left {
		current = settingsScreen
		return
	}
	if right {
		current = settingsSettingAlarmToggleTypeScreen
		return
	}
	// Update the screen.
	line1 = "Settings:"
	line2 = "Alarm Type~"
	return
}

type settingsSettingAlarmToggleType struct{}

func (screen *settingsSettingAlarmToggleType) Update(s State, up, down, left, right bool, AlarmTimeHour int) (state State, current Screen, line1, line2 string) {
	// Set the output state.
	state = s
	// Keep on the same screen.
	current = screen

	// Handle events.

	if up && state.AlarmType < 2 {
		state.AlarmType++
	}
	if down && state.AlarmType > 0 {
		state.AlarmType--
	}

	if left {
		current = settingsSetAlarmTypeScreen
		return
	}
	// Update the screen.
	line1 = "Alarm Type:"
	line2 = alarmtypedecoder(state.AlarmType)
	return
}

// Screens of the system.

// HomeScreen is the first screen that gets loaded.
var alarmStateScreen = &alarmState{}
var calendarScreen Screen = &calendar{}
var clockScreen Screen = &clock{}
var HomeScreen Screen = &home{}
var settingsScreen = &settings{}
var settingsSetTimeScreen = &settingsSetTime{}
var settingsSetAlarmTimeScreen = &settingsSetAlarmTime{}
var settingsSetAlarmToggleScreen = &settingsSetAlarmToggle{}
var settingsSetAlarmTypeScreen = &settingsSetAlarmToggleType{}

var settingsSettingTimeHourScreen = &settingsSettingTimeHour{}
var settingsSettingTimeMinScreen = &settingsSettingTimeMin{}
var settingsSettingAlarmTimeHourScreen = &settingsSettingAlarmTimeHour{}
var settingsSettingAlarmTimeMinScreen = &settingsSettingAlarmTimeMin{}
var settingsSettingAlarmToggleScreen = &settingsSettingAlarmToggle{}
var settingsSettingAlarmToggleTypeScreen = &settingsSettingAlarmToggleType{}
