package main

import (
	"runtime"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 800
	screenHeight = 600

	shipMaxSpeed    = 4
	shipAccel       = 0.25
	shipDecel       = 0.15
	shipRotateSpeed = 3

	shootCooldown = 250 * time.Millisecond

	laserSpeed = 8
)

var res *Resources
var lasers []*Laser

func main() {
	runtime.LockOSThread()

	res = NewResources()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Rectangle", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	player := NewPlayer([5]sf.KeyCode{sf.KeyUp, sf.KeyDown, sf.KeyLeft, sf.KeyRight, sf.KeySpace}, sf.Vector2f{screenWidth / 2, screenHeight / 2})

	var dt float32
	for window.IsOpen() {
		start := time.Now()

		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventClosed:
				window.Close()
			}
		}

		player.Update(dt)

		for _, l := range lasers {
			l.Update(dt)
		}

		window.Clear(sf.ColorWhite)

		window.Draw(player)

		for _, l := range lasers {
			window.Draw(l)
		}

		window.Display()

		elapsed := time.Since(start)
		dt = float32(elapsed) / float32(time.Second)
	}
}
