package main

import (
	"math"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Player struct {
	*sf.Sprite

	speed    float32
	keys     [5]sf.KeyCode
	canShoot bool
}

func NewPlayer(keys [5]sf.KeyCode, pos sf.Vector2f) *Player {
	player := new(Player)
	player.Sprite = sf.NewSprite(res.images["playerShip1_blue.png"])

	player.SetPosition(pos)

	player.keys = keys
	player.canShoot = true

	size := player.GetGlobalBounds()
	player.SetOrigin(sf.Vector2f{size.Width / 2, size.Height / 2})

	return player
}

func (p *Player) Shoot() {
	if p.canShoot {
		laser := NewLaser(p.GetPosition(), p.GetRotation(), laserSpeed)
		lasers = append(lasers, laser)

		p.canShoot = false

		go func() {
			time.Sleep(shootCooldown)
			p.canShoot = true
		}()
	}
}

func (p *Player) Update(dt float32) {
	if sf.KeyboardIsKeyPressed(p.keys[0]) {
		if p.speed < shipMaxSpeed {
			p.speed += shipAccel
		}
	} else {
		if p.speed > 0 {
			p.speed -= shipDecel
		}
	}

	if sf.KeyboardIsKeyPressed(p.keys[2]) {
		p.Rotate(-shipRotateSpeed * dt * 60)
	}
	if sf.KeyboardIsKeyPressed(p.keys[3]) {
		p.Rotate(shipRotateSpeed * dt * 60)
	}

	if sf.KeyboardIsKeyPressed(p.keys[4]) {
		p.Shoot()
	}

	if p.speed < 0 {
		p.speed = 0
	}
	if p.speed > shipMaxSpeed {
		p.speed = shipMaxSpeed
	}

	angle := p.GetRotation() - 90
	angleRad := angle * math.Pi / 180

	vx := p.speed * float32(math.Cos(float64(angleRad)))
	vy := p.speed * float32(math.Sin(float64(angleRad)))

	p.Move(sf.Vector2f{vx * dt * 60, vy * dt * 60})
}
