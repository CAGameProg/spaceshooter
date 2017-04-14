# Asteroids

Create a new file called `asteroid.go` which will contain the code for the `Asteroid` object.

Asteroids should fly around the screen while rotating and also wrapping to the other side of the screen when they are no longer visible.

First you should create the `Asteroid` struct:

```go
type Asteroid struct {
    // Fields...
}
```

Think about what kind of data an asteroid should store (the `Asteroid` is probably
quite similar to the `Laser` struct).

Remember that unlike lasers and players,
asteroids are constantly spinning but still moving in the same direction. This means
that you need to store an additional angle (`Rotation` in the sprite will keep track
of the visual angle and you probably want to store another variable in the struct to
keep track of the move angle).

Also, the asteroid's rotation speed is not constant because it should be different
for every asteroid, so it needs to be stored in the struct.

# Constructor

Write a `NewAsteroid` function which will create a new asteroid:

```go
func NewAsteroid(pos sf.Vector2f, rotationSpeed, rotationAngle, speed float32) *Asteroid {
    // ...
}
```

This function should create a new asteroid object with the correct qualities.

**Note**: For the asteroid images, you must download the new images placed in the [asteroids](../assets/images/asteroids) folder. Place the `asteroids` folder in `assets/images`.

You may choose one of the images, or look in the Bonus section at the bottom a guide
on how to randomly select an image.

# Update method

Now create an update method for the asteroid. It should move the asteroid along its
move angle, and rotate the asteroid according to its rotation speed.

Here is the type signature:

```go
func (a *Asteroid) Update(dt float32) {
    // ...
}
```

Remember to use `dt` when rotating as well (you may want to look in `player.go`).

# Wrap method

In addition, you should write a `Wrap` method which will wrap the asteroid to the
other side of the screen when it moves off the screen:

```go
func (a *Asteroid) Wrap() {
    // ...
}
```

If you are struggling, you can look at the hint below.

<details>
  <summary>Click to see hint</summary>
  ```
  if x position < 0 - width/2:
      set x position to screenWidth + width/2
  if x position > screenWidth + width/2:
      set x position to 0 - width/2

  if y position < 0 - height/2:
      set y position to screenHeight + height/2
  if y position > screenHeight + height/2:
      set y position to 0 - height/2
    ```

  You may want to use `GetPosition`, `SetPosition` and `GetGlobalBounds`.
</details>

**Note**: remember to call `a.Wrap()` in your `Update` method so it gets checked
every frame.

---

Try to test your program by creating an asteroid at a certain position and calling
`asteroids.Update(dt)` and `window.Draw(asteroid)` in the main loop.

# Bonus: Random asteroid images

You might also want to have the asteroids each use a random image.

You can use:

```go
num := rand.Intn(4) + 1
```

to get a random number between 1 and 4.

Then use

```go
res.images["asteroid"+strconv.Itoa(num)+".png"]
```

to get a random texture from the asset loader.

**Note**: remember to put the line:

```go
rand.Seed(time.Now().UTC().UnixNano())
```

at the beginning of the `main` function.

# Bonus: Spawning asteroids

Make a constant called `numAsteroids` (I recommend 5 to 10).

In your main function, create a bunch of asteroids with a for loop and put them into an array which is updated every frame.

You'll want to give the asteroids random speeds, rotations, rotationSpeeds, and positions.

Use this function to generate a random float:

```go
func random(min, max float32) float32 {
    return rand.Float32()*(max-min) + min
}
```

**Note**: remember to put the line:

```go
rand.Seed(time.Now().UTC().UnixNano())
```

at the beginning of the `main` function.
