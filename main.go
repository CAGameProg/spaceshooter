package main

import (
	"math/rand"
	"runtime"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 1200
	screenHeight = 800

	shipMaxSpeed    = 4
	shipAccel       = 0.25
	shipDecel       = 0.15
	shipRotateSpeed = 3

	shootCooldown = 250 * time.Millisecond

	laserSpeed = 8

	numAsteroids = 6
)

func random(min, max float32) float32 {
	return rand.Float32()*(max-min) + min
}

var res *Resources
var lasers []*Laser
var asteroids []*Asteroid

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	runtime.LockOSThread()

	res = NewResources()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Rectangle", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	player := NewPlayer([5]sf.KeyCode{sf.KeyUp, sf.KeyDown, sf.KeyLeft, sf.KeyRight, sf.KeySpace}, sf.Vector2f{screenWidth / 2, screenHeight / 2})

	for i := 0; i < numAsteroids; i++ {
		pos := sf.Vector2f{random(0, screenWidth), random(0, screenHeight)}
		speed := random(3, 7)
		rotationSpeed := random(2, 5)
		rotationAngle := random(0, 360)

		asteroid := NewAsteroid(pos, rotationSpeed, rotationAngle, speed)
		asteroids = append(asteroids, asteroid)
	}

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
		for _, a := range asteroids {
			a.Update(dt)
		}

		for _, l := range lasers {
			for _, a := range asteroids {
				if Intersects(l.Sprite, a.Sprite) {
					a.dead = true
					l.dead = true

					soundBuffer := res.sounds["sfx_explosion.wav"]
					sound := sf.NewSound(soundBuffer)
					sound.Play()
				}
			}
		}

		window.Clear(sf.ColorBlack)

		window.Draw(player)

		for _, l := range lasers {
			window.Draw(l)
		}
		for _, a := range asteroids {
			window.Draw(a)
		}

		window.Display()

		var tempLasers []*Laser
		for _, l := range lasers {
			if !l.dead {
				tempLasers = append(tempLasers, l)
			}
		}
		lasers = tempLasers

		var tempAsteroids []*Asteroid
		for _, a := range asteroids {
			if !a.dead {
				tempAsteroids = append(tempAsteroids, a)
			}
		}
		asteroids = tempAsteroids

		elapsed := time.Since(start)
		dt = float32(elapsed) / float32(time.Second)
	}
}
