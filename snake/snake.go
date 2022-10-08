package snake

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"time"
)

var ErrUnauthorizedMove = errors.New("unauthorized move")

const gridSize = 10

type Snake struct {
	parts    []*Part
	cherries []*Cherry

	gridSizeX, gridSizeY, visualSize int // Size of the grid in which the snake is evolving
	board                            *Board

	directionX int // -1 (go to left), 0 (no move), 1 (go to right)
	directionY int // -1 (go to up), 0 (no move), 1 (go to down)

	tickMove   int
	tickCherry int
}

func NewSnake(layoutWidth, layoutHeight int) *Snake {
	snake := new(Snake)
	rand.Seed(time.Now().UnixNano())

	// New board
	snake.board = NewBoard(5, 15, float64(layoutWidth-5), float64(layoutHeight-5))

	fmt.Printf("Width: %d, Heigth: %d\n", snake.board.Width(), snake.board.Height())

	// Store important data
	snake.gridSizeX = snake.board.Width() / gridSize
	snake.gridSizeY = snake.board.Height() / gridSize
	snake.visualSize = gridSize
	snake.tickMove = 0
	snake.tickCherry = 0

	// Add the head of the snake at random position on the grid
	snake.parts = make([]*Part, 1)
	snake.parts[0] = NewPart(rand.Intn(snake.gridSizeX), rand.Intn(snake.gridSizeY), snake.visualSize, snake.board)

	// Create an empty cherry vector
	snake.cherries = make([]*Cherry, 0)

	// Based on the position of the head, define the directions
	snake.initDirection()

	return snake
}

func (s *Snake) Update() error {
	// Every X ticks, apply the move if it is allowed
	if s.tickMove == 5 {
		if s.isMoveAllowed() {
			if s.isEatingCherry() {
				s.eatCherryAndIncreaseSnake()
			} else {
				s.applyMove()
			}
		} else {
			return ErrUnauthorizedMove
		}
		s.tickMove = 0
	}

	// Every X ticks, add a berry
	if s.tickCherry == 120 {
		s.createCherry()
		s.tickCherry = 0
	}

	s.tickCherry++
	s.tickMove++

	return nil
}

func (s *Snake) Score() int {
	return len(s.parts)
}

func (s *Snake) Draw(screen *ebiten.Image) {
	s.board.Draw(screen)

	for _, part := range s.parts {
		part.Draw(screen)
	}

	for _, cherry := range s.cherries {
		cherry.Draw(screen)
	}
}

func (s *Snake) applyMove() {
	// Each part is taking the place of the previous part
	for i := len(s.parts) - 1; i > 0; i-- {
		s.parts[i].X = s.parts[i-1].X
		s.parts[i].Y = s.parts[i-1].Y
	}

	// The head is taking the future position
	s.parts[0].X = s.parts[0].X + s.directionX
	s.parts[0].Y = s.parts[0].Y + s.directionY
}

func (s *Snake) isMoveAllowed() bool {
	// Compute future position of the head
	futureX, futureY := s.getNextPosition()

	// Rule #1 : The head cannot go outside the limits of the grid
	if futureX < 0 || futureX >= s.gridSizeX {
		return false
	}
	if futureY < 0 || futureY >= s.gridSizeY {
		return false
	}

	// Rule #2 : The head cannot be on the same position of
	// a future position of a part of the snake (so, we ignore the
	// tail of the snake

	// If only 2 parts of the snake, no issues
	if len(s.parts) <= 2 {
		return true
	}

	// Check each part if longer...
	for _, part := range s.parts[1 : len(s.parts)-2] {
		if futureX == part.X && futureY == part.Y {
			return false
		}
	}

	return true
}

func (s *Snake) isEatingCherry() bool {
	// Compute future position of the head
	futureX, futureY := s.getNextPosition()

	// Check if there is a cherry
	for _, cherry := range s.cherries {
		if futureX == cherry.X && futureY == cherry.Y {
			return true
		}
	}

	return false
}

func (s *Snake) initDirection() {
	// Randomly go horizontally or vertically
	if rand.Intn(2) == 0 {
		// Horizontally
		s.directionY = 0
		if float32(s.parts[0].X)/float32(s.gridSizeX) < 0.5 {
			// We are on the left side of the board, go to right
			s.directionX = +1
		} else {
			// We are on the right side of the board, go to left
			s.directionX = -1
		}
	} else {
		// Vertically
		s.directionX = 0
		if float32(s.parts[0].Y)/float32(s.gridSizeY) < 0.5 {
			// We are on the upper side of the board, go down
			s.directionY = +1
		} else {
			// We are on the lower side of the board, go up
			s.directionY = -1
		}
	}
}

func (s *Snake) ApplyDirection(keys []ebiten.Key) {
	// No key pressed
	if len(keys) == 0 {
		return
	}

	// More than one key pressed -> ignore
	if len(keys) > 1 {
		return
	}

	switch keys[0] {
	case ebiten.KeyArrowRight:
		s.directionX = +1
		s.directionY = 0
	case ebiten.KeyArrowLeft:
		s.directionX = -1
		s.directionY = 0
	case ebiten.KeyArrowUp:
		s.directionX = 0
		s.directionY = -1
	case ebiten.KeyArrowDown:
		s.directionX = 0
		s.directionY = +1
	}
}

func (s *Snake) createCherry() {
	var futureX, futureY int
	isCollision := true

	for isCollision {
		// Compute new position
		futureX = rand.Intn(s.gridSizeX)
		futureY = rand.Intn(s.gridSizeY)
		isCollision = false

		// Check if collision with snake
		for _, part := range s.parts {
			if part.X == futureX && part.Y == futureY {
				isCollision = true
			}
		}

		// Check if collision with other cherries
		for _, cherry := range s.cherries {
			if cherry.X == futureX && cherry.Y == futureY {
				isCollision = true
			}
		}
	}

	newCherry := NewCherry(futureX, futureY, s.board)
	s.cherries = append(s.cherries, newCherry)
}

func (s *Snake) eatCherryAndIncreaseSnake() {
	futureX, futureY := s.getNextPosition()
	s.removeCherryAt(futureX, futureY)

	lastX := s.parts[len(s.parts)-1].X
	lastY := s.parts[len(s.parts)-1].Y
	s.applyMove()

	newTail := NewPart(lastX, lastY, s.visualSize, s.board)
	s.parts = append(s.parts, newTail)
}

func (s *Snake) removeCherryAt(x int, y int) {
	idxToRemove := -1
	for i := range s.cherries {
		if s.cherries[i].X == x && s.cherries[i].Y == y {
			idxToRemove = i
			break
		}
	}

	if idxToRemove >= 0 {
		s.cherries = append(s.cherries[:idxToRemove], s.cherries[idxToRemove+1:]...)
	}
}

func (s *Snake) getNextPosition() (x, y int) {
	x = s.parts[0].X + s.directionX
	y = s.parts[0].Y + s.directionY
	return
}
