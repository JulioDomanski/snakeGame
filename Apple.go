package main

import "math/rand/v2"

type Apple struct {
	X int
	Y int
}

func (a *Apple) SpawnFood(screenX int, screenY int) {
	a.X = rand.IntN((screenWidth-10)/gridSize) * gridSize
	a.Y = rand.IntN((screenHeight-10)/gridSize) * gridSize
}
