package main

import (
	"math"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Laser struct {
	*sf.Sprite

	speed float32
}

func NewLaser(pos sf.Vector2f, rotation float32, speed float32) *Laser {
	laser := new(Laser)
	laser.Sprite = sf.NewSprite(res.images["laserBlue.png"])

	laser.SetPosition(pos)

	size := laser.GetGlobalBounds()
	laser.SetOrigin(sf.Vector2f{size.Width / 2, size.Height / 2})

	laser.speed = speed
	laser.SetRotation(rotation)

	return laser
}

func (l *Laser) Update(dt float32) {
	angle := l.GetRotation() - 90
	angleRad := angle * math.Pi / 180

	vx := l.speed * float32(math.Cos(float64(angleRad)))
	vy := l.speed * float32(math.Sin(float64(angleRad)))

	l.Move(sf.Vector2f{vx * dt * 60, vy * dt * 60})
}
