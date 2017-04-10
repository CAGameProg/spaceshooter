# Laser and Projectiles

Create a new file called `laser.go` and write some code for a laser object.

First you should create the `Laser` struct:

```go
type Laser struct {
    // Fields...
}
```

Think about what fields you may want to store in a laser and what structs you might want to embed (the `Laser` is very similar to the `Player` struct).

# Constructor

You'll also want to write a `NewLaser` function which will allow you to create a new laser:

```go
func NewLaser(pos sf.Vector2f, rotation float32, speed float32) *Laser {
    // ...
}
```

This function should create a new laser object at `pos` and rotation `rotation` and with speed `speed`.
You should also make sure to set the origin of the laser sprite in the center of the image rather than the top left corner.

I recommend using the `laserBlue.png` file for the image, although you can use a different color if you would like (`laserGreen.png` or `laserRed.png`).

# Update method

Now create an update method for the laser which will move the laser forward along its angle of rotation at the speed set in the `NewLaser` function.

Here is the type signature of the method:

```go
func (l *Laser) Update(dt float32) {
    // ...
}
```

---

If you would like to test your program, you can create a laser in your main function and call
`laser.Update(dt)` and `window.Draw(laser)`. Next class we will discuss how to spawn lasers by pressing
a key and how to have an arbitrary number of lasers on the screen at once.
