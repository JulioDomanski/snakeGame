package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	gridSize     = 5
)

type Snake struct {
	Body []Point
	DirX int
	DirY int
}

func (s *Snake) Grow() {
	head := s.Body[0]
	newHead := Point{head.X + s.DirX*gridSize, head.Y + s.DirY*gridSize}
	s.Body = append([]Point{newHead}, s.Body...)
}

func (s *Snake) Move() {

	head := s.Body[0]
	newHead := Point{head.X + s.DirX*gridSize, head.Y + s.DirY*gridSize}

	s.Body = append([]Point{newHead}, s.Body[:len(s.Body)-1]...)

}

type Point struct {
	X, Y int
}

type Game struct {
	Snake *Snake
	Apple *Apple
}

func (g *Game) Update() error {
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
	fmt.Println(g.Snake.Body[0].X, g.Snake.Body[0].Y)
	fmt.Println(g.Apple.X, g.Apple.Y)
	if g.Snake.Body[0].X >= g.Apple.X && g.Snake.Body[0].X < (g.Apple.X+11) &&
		g.Snake.Body[0].Y >= g.Apple.Y && g.Snake.Body[0].Y < (g.Apple.Y+11) {
		g.Snake.Grow()
	}
	g.Snake.Move()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	snakeColor := color.RGBA{255, 255, 255, 255}
	appleColor := color.RGBA{255, 0, 0, 250}

	for _, segment := range g.Snake.Body {
		screen.Set(segment.X, segment.Y, snakeColor)
	}
	appleSize := 10
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
		Snake: &Snake{
			Body: []Point{{X: screenWidth / 2, Y: screenHeight / 2}},
		},
		Apple: &Apple{
			X: 500,
			Y: 400,
		},
	}
	// Inicia o jogo
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
