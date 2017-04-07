package main

import (
	"runtime"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var res *Resources

func main() {
	runtime.LockOSThread()

	res = NewResources()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Rectangle", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	player := NewPlayer(sf.Vector2f{screenWidth / 2, screenHeight / 2})

	for window.IsOpen() {
		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventClosed:
				window.Close()
			}
		}

		window.Clear(sf.ColorWhite)

		window.Draw(player)

		window.Display()
	}
}
