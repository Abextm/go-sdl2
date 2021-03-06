// This program creates an SDL window and renderer and takes in events using sdl.PeepEvents
package main

import (
	"fmt"
	"github.com/abextm/go-sdl2/sdl"
	"os"
)

var winTitle string = "Go-SDL2 TestWaitEvent"
var winWidth, winHeight int = 800, 600

func main() {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var event sdl.Event
	var running bool

	window = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if window == nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", sdl.GetError())
		os.Exit(1)
	}

	renderer = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if renderer == nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", sdl.GetError())
		os.Exit(2)
	}

	var peepArray []sdl.Event = make([]sdl.Event, 5)

	running = true
	for running {
		sdl.PumpEvents()
		numEventsRetrieved := sdl.PeepEvents(peepArray, sdl.PEEKEVENT, sdl.FIRSTEVENT, sdl.LASTEVENT)
		if numEventsRetrieved < 0 {
			fmt.Printf("PeepEvents error: %s\n", sdl.GetError())
		} else {
			for i := 0; i < numEventsRetrieved; i++ {
				fmt.Printf("Event Peeked Value: %v\n", peepArray[i]) // primitive printing of event
			}
		}

		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
			case *sdl.MouseWheelEvent:
				fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y)
			case *sdl.KeyUpEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			}
		}
		sdl.Delay(1000 / 30)
	}

	renderer.Destroy()
	window.Destroy()
}
