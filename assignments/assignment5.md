# Explosions

We have already created the explosion struct in class. But now we need to make sure that every explosion is displayed every frame.

Create a global `explosions` array which will hold the explosion objects. You should put the declaration in `main.go` next to the `lasers`, `asteroids`... arrays.

Now create a function in `util.go` called `SpawnExplosion` which will create a new explosion and place it in the array. `SpawnExplosion` should create the explosion at
a given location. Here is the function definition:

```go
func SpawnExplosion(pos sf.Vector2f) {
    // Your code...
}
```

This function should:

* Create a new explosion using the `NewExplosion` function we wrote in class (located in `explosion.go`).
* Place the newly created explosion object in the `explosions` array. You should have created this array as a global array so it should be accessible anywhere.

---

Now in `main.go` between `window.Clear(...)` and `window.Display`, you should draw all the explosions by looping through every object in the `explosions` array and
calling `window.Draw(...)` on it.

---

After `window.Display` create some logic similar to the ones for `lasers` and `asteroids` which should remove the dead explosions from the `explosions` array.

---

Now you probably want to spawn the explosions when a laser collides with an asteroid. We made a double for loop to check for intersections in `main.go`. Inside this for loop,
if the asteroid `a` is intersecting with the laser `l`, you should call `SpawnExplosion(a.GetPosition())` to spawn an explosion at the asteroid's position.
