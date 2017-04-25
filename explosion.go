package main

import (
	"strconv"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Explosion struct {
	*sf.Sprite

	curFrame int
	dead     bool
}

func NewExplosion(pos sf.Vector2f) *Explosion {
	e := new(Explosion)
	e.Sprite = sf.NewSprite(res.images["explosion_0.png"])

	e.SetPosition(pos)

	size := e.GetGlobalBounds()
	e.SetOrigin(sf.Vector2f{size.Width / 2, size.Height / 2})

	go e.Animate()

	return e
}

func (e *Explosion) Animate() {
	for i := 0; i < 24; i++ {
		e.Sprite.SetTexture(res.images["explosion_"+strconv.Itoa(i)+".png"], false)
		time.Sleep(20 * time.Millisecond)
	}
	e.dead = true
}
