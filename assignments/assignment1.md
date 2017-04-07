# 1. Set the player's origin

One thing you might notice is that the top left corner of the player is
drawn at the coordinates of the player. We probably want to change this
so that the image is centered at the coordinates.

To do this, you must set the player's origin to half the image's width
and half the image's height.

Add some code to the `NewPlayer` function in `player.go` that will
correctly set the origin so that the texture is centered.

To set the origin, you should use the `SetOrigin` method:

* `SetOrigin(vec sf.Vector2f)`: this method is attached to the Sprite struct but can be called on a Player because Player embeds Sprite.

To get the width and height of the sprite, you can use the following method:

* `GetGlobalBounds() sf.Rectf`: this method is attached to the Sprite struct but can be called on a Player because Player embeds Sprite.

This returns a `Rectf` object. You can use `rect.Width` and `rect.Height` to
access the width and height respectively.

# 2. Write a simple Update method

Next you can try to write an `Update` method for your player struct. For now
the `Update` method should check if any keys are pressed, and if so, move
the sprite up, down, left, or right accordingly. If no key is pressed don't
move the player.

You should define the `Update` method like so:

```go
func (p *Player) Update() {
    // Your code here...
}
```

Remember that to check if a key is pressed, you should use the following function:

* `sf.KeyboardIsKeyPressed(key sf.KeyCode) bool`: Remember that Player has a field called `keys` which will give you the keys to check. `keys[0]` = up, `keys[1]` = down, `keys[2]` = left, `keys[3]` = right.

For example, to check if the up arrow is pressed:

```go
if sf.KeyboardIsKeyPressed(p.keys[0]) {
    // Up arrow is currently pressed...
}
```

Then to move the player, you should call 

* `Move(amount sf.Vector2f)`: this method is attached to the Sprite struct but can be called on a Player because Player embeds Sprite.

For example, to move the Player down 2 pixels (remember that larger coordinates is farther down the screen because `(0,0)` is the top left corner):

```go
p.Move(sf.Vector2f{0, 2})
```

Remember to call `player.Update()` in the loop in `main.go`. I recommend you call the method right before `window.Clear` (~line 34).
