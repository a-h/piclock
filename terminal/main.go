package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gdamore/tcell"
	"github.com/headblockhead/alarmclock"
)

type keypressEvent struct {
	up, down, left, right bool
}

func main() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sigs
		cancel()
	}()

	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	s.SetStyle(lcdStyle.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	s.Clear()

	events := make(chan keypressEvent)
	state := alarmclock.NewState()
	screen := alarmclock.HomeScreen
	state, screen, line1, line2 := screen.Update(state, false, false, false, false)
	render(s, line1, line2)

	go func() {
		for {
			var kpe keypressEvent

			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyCtrlC:
					cancel()
				case tcell.KeyUp:
					kpe.up = true
				case tcell.KeyDown:
					kpe.down = true
				case tcell.KeyLeft:
					kpe.left = true
				case tcell.KeyRight:
					kpe.right = true
				case tcell.KeyEsc:
					cancel()
				case tcell.KeyCtrlL:
					s.Sync()
				}
			case *tcell.EventResize:
				s.Sync()
			}
			events <- kpe
			if ctx.Err() != nil {
				break
			}
		}
	}()

Loop:
	for {
		var kpe keypressEvent
		// Wait for a keypress, but render at least once per second.
		select {
		case _ = <-time.After(time.Second):
		case kpe = <-events:
		case <-ctx.Done():
			break Loop
		}
		// Check for updates.
		var newLine1, newLine2 string
		var newScreen alarmclock.Screen
		state, newScreen, newLine1, newLine2 = screen.Update(state, kpe.up, kpe.down, kpe.left, kpe.right)
		screenUpdated := screen != newScreen
		screen = newScreen
		// If the screen has changed, render the new screen.
		if screenUpdated {
			state, newScreen, newLine1, newLine2 = screen.Update(state, kpe.up, kpe.down, kpe.left, kpe.right)
		}
		// Only update when the screen needs it.
		if line1 != newLine1 || line2 != newLine2 {
			line1 = newLine1
			line2 = newLine2
			render(s, line1, line2)
		}
	}

	s.Fini()
}

var lcdStyle = tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)

func render(s tcell.Screen, line1, line2 string) {
	// Clear everything.
	for y := 0; y < 4; y++ {
		for x := 0; x < 18; x++ {
			s.SetContent(x, y, ' ', nil, lcdStyle)
		}
	}

	for i := 0; i < len(line1); i++ {
		s.SetContent(i+1, 1, rune(line1[i]), nil, lcdStyle)
	}
	for i := 0; i < len(line2); i++ {
		s.SetContent(i+1, 2, rune(line2[i]), nil, lcdStyle)
	}

	s.Show()
}
