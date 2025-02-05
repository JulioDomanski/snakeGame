package main

type Snake struct {
	Body []Point
	DirX int
	DirY int
}

type Point struct {
	X, Y int
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
