package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/stianeikeland/go-rpio"

	"github.com/a-h/character"
	"github.com/a-h/piclock"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

var display *character.Display

func initialiseLCD() {
	// Use the periph library.
	_, err := host.Init()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	// Open up first i2c channel.
	// You'll need to enable i2c for your Raspberry Pi in
	// https://www.raspberrypi.org/documentation/configuration/raspi-config.md
	bus, err := i2creg.Open("")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	// The default address for the i2c backpack is 0x27.
	dev := &i2c.Dev{
		Bus:  bus,
		Addr: 0x27,
	}

	// Create a 2 line display.
	display = character.NewDisplay(dev, false)
}

func main() {
	initialiseLCD()

	// Handle ctrl+c.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	rpio.Open()
	defer rpio.Close()
	up := Debounce(rpio.Pin(18))
	down := Debounce(rpio.Pin(15))
	left := Debounce(rpio.Pin(23))
	right := Debounce(rpio.Pin(24))

	state := piclock.NewState()
	screen := piclock.HomeScreen
	var line1, line2 string

	for {
		var upState, downState, leftState, rightState bool
		if buttonState, changed := up(); changed && buttonState == rpio.Low {
			fmt.Println("UP!")
			upState = true
		}
		if buttonState, changed := down(); changed && buttonState == rpio.Low {
			fmt.Println("DOWN!")
			downState = true
		}
		if buttonState, changed := left(); changed && buttonState == rpio.Low {
			fmt.Println("LEFT!")
			leftState = true
		}
		if buttonState, changed := right(); changed && buttonState == rpio.Low {
			fmt.Println("RIGHT!")
			rightState = true
		}
		// Check for updates.
		var newLine1, newLine2 string
		var newScreen piclock.Screen
		state, newScreen, newLine1, newLine2 = screen.Update(state, upState, downState, leftState, rightState)
		screenUpdated := screen != newScreen
		screen = newScreen
		// If the screen has changed, render the new screen.
		if screenUpdated {
			state, newScreen, newLine1, newLine2 = screen.Update(state, upState, downState, leftState, rightState)
		}
		// Only update when the screen needs it.
		if line1 != newLine1 || line2 != newLine2 {
			line1 = newLine1
			line2 = newLine2
			display.Goto(0, 0)
			display.Print(line1 + "                ")
			display.Goto(1, 0)
			display.Print(line2 + "                ")
		}
		select {
		case <-sigs:
			display.Clear()
			return
		case <-time.After(time.Millisecond):
			break
		}
	}
}

// Debounce a pin.
func Debounce(pin rpio.Pin) func() (s rpio.State, updated bool) {
	pin.PullUp()
	lastChange := time.Now()
	state := pin.Read()
	return func() (rpio.State, bool) {
		if time.Now().Before(lastChange.Add(time.Millisecond * 10)) {
			return state, false
		}
		prev := state
		state = pin.Read()
		if prev != state {
			lastChange = time.Now()
		}
		return state, prev != state
	}
}
