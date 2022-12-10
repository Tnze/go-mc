# Go-MC


Minecraft for "rodents"

## Badges
![Downloads](https://img.shields.io/github/downloads/Edouard127/go-mc/total)
![Activity](https://img.shields.io/github/commit-activity/w/Edouard127/go-mc)
![Size](https://img.shields.io/github/languages/code-size/Edouard127/go-mc)

## Features

- Ray-Tracing

## Todo

- [ ] Game physics
- [ ] Inventory transactions
- [ ] Movements
- [ ] Chat
- [ ] A* pathfinding

## Documentation

### Ray-Tracing

You can cast a ray in the world to get the block you are looking at.
```go
start := c.Player.GetEyePos()
end := maths.ProjectPosition(c.Player.Rotation, 5, 1.62) // Relative to the player's eye position
result, err := c.World.RayTrace(start, start.Add(end))
if err != nil {
fmt.Println(err)
}
fmt.Println(result.String())
```
