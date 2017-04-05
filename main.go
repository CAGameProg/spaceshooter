package main

import (
	"fmt"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

func main() {
	res := NewResources() // You will write this function

	// Now you can use textures and assets to play a sound for example:
	sound := sf.NewSound(res.sounds["sfx_laser1.ogg"])
	sound.Play()

	// Or you could make a new sprite from an image:
	fmt.Println(res.images["playerShip1_blue.png"])

	time.Sleep(2 * time.Second)
}
