package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 680
	screenHeight = 340
	gridSize     = 2
)

type Game struct {
	Snake *Snake
	Apple *Apple
	Point int
}

func (g *Game) Update() error {

	if g.Snake.Body[0].X == 680 ||
		g.Snake.Body[0].X == 0 ||
		g.Snake.Body[0].Y == 0 ||
		g.Snake.Body[0].Y == 340 {
		fmt.Println("cabo")
		return errors.New("game ended by player")
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.Snake.DirX, g.Snake.DirY = 0, -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.Snake.DirX, g.Snake.DirY = 0, 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.Snake.DirX, g.Snake.DirY = -1, 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.Snake.DirX, g.Snake.DirY = 1, 0
	}
	if g.Snake.Body[0].X < g.Apple.X+5 && g.Snake.Body[0].X+5 > g.Apple.X &&
		g.Snake.Body[0].Y < g.Apple.Y+5 && g.Snake.Body[0].Y+5 > g.Apple.Y {
		g.Snake.Grow()
		g.Apple.SpawnFood(screenWidth, screenHeight)
		g.Point += 10
	}
	fmt.Println(g.Snake.Body[0].X, g.Snake.Body[0].Y)
	g.Snake.Move()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	text.Draw(screen, fmt.Sprintf("Score: %d", g.Point), basicfont.Face7x13, 10, 20, color.RGBA{255, 255, 255, 255})
	snakeColor := color.RGBA{255, 255, 255, 255}
	appleColor := color.RGBA{255, 0, 0, 250}

	segmentSize := 5

	for _, segment := range g.Snake.Body {
		snakeRect := ebiten.NewImage(segmentSize, segmentSize)
		snakeRect.Fill(snakeColor)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(segment.X), float64(segment.Y))
		screen.DrawImage(snakeRect, op)
	}
	appleSize := 5
	appleRect := ebiten.NewImage(appleSize, appleSize)
	appleRect.Fill(appleColor)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.Apple.X), float64(g.Apple.Y))
	screen.DrawImage(appleRect, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake")

	game := &Game{
		Point: 0,
		Snake: &Snake{
			Body: []Point{{X: screenWidth / 2, Y: screenHeight / 2}},
		},
		Apple: &Apple{
			X: 320,
			Y: 200,
		},
	}
	// Inicia o jogo
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
