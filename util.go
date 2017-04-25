package main

import sf "github.com/zyedidia/sfml/v2.3/sfml"

func Intersects(s1, s2 *sf.Sprite) bool {
	bounds1 := s1.GetGlobalBounds()
	bounds2 := s2.GetGlobalBounds()

	didCollide, _ := bounds1.Intersects(bounds2)
	return didCollide
}

func Wrap(s *sf.Sprite) {
	pos := s.GetPosition()
	size := s.GetGlobalBounds()

	xPos := pos.X
	yPos := pos.Y

	if xPos < 0-size.Width/2 {
		pos.X = screenWidth + size.Width/2
	}
	if xPos > screenWidth+size.Width/2 {
		pos.X = 0 - size.Width/2
	}
	if yPos < 0-size.Height/2 {
		pos.Y = screenHeight + size.Height/2
	}
	if yPos > screenHeight+size.Height/2 {
		pos.Y = 0 - size.Height/2
	}

	s.SetPosition(pos)
}
