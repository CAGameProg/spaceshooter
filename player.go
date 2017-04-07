package main

import sf "github.com/zyedidia/sfml/v2.3/sfml"

type Player struct {
	*sf.Sprite

	speed float32
	keys  [5]sf.KeyCode
}

func NewPlayer(pos sf.Vector2f) *Player {
	player := new(Player)
	player.Sprite = sf.NewSprite(res.images["playerShip1_blue.png"])

	player.SetPosition(pos)

	player.keys = [5]sf.KeyCode{sf.KeyUp, sf.KeyDown, sf.KeyLeft, sf.KeyRight, sf.KeySpace}

	return player
}
