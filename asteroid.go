package main

import (
	"math"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Asteroid struct {
	*sf.Sprite

	moveAngle     float32
	rotationSpeed float32

	speed float32

	dead bool
}

func NewAsteroid(pos sf.Vector2f, rotationSpeed, rotationAngle, speed float32) *Asteroid {
	asteroid := new(Asteroid)

	asteroid.Sprite = sf.NewSprite(res.images["asteroid1.png"])

	bounds := asteroid.GetGlobalBounds()
	asteroid.SetOrigin(sf.Vector2f{bounds.Width / 2, bounds.Height / 2})

	asteroid.speed = speed
	asteroid.SetRotation(rotationAngle)
	asteroid.moveAngle = rotationAngle
	asteroid.rotationSpeed = rotationSpeed

	asteroid.SetPosition(pos)

	return asteroid
}

func (a *Asteroid) Update(dt float32) {
	a.Rotate(a.rotationSpeed * 60 * dt)

	angle := a.moveAngle
	angleRad := angle * math.Pi / 180

	vx := a.speed * float32(math.Cos(float64(angleRad)))
	vy := a.speed * float32(math.Sin(float64(angleRad)))

	a.Move(sf.Vector2f{vx * dt * 60, vy * dt * 60})

	Wrap(a.Sprite)
}
