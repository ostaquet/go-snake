package snake

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var ErrUnauthorizedMove = errors.New("unauthorized move")

const gridSize = 10

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Snake struct {
	parts    []*Part
	cherries []*Cherry

	gridSizeX, gridSizeY, visualSize int // Size of the grid in which the snake is evolving
	board                            *Board
	mplusSmallFont                   font.Face

	direction Direction

	tickMove   int
	tickCherry int
}

func NewSnake(layoutWidth, layoutHeight int) *Snake {
	snake := new(Snake)
	rand.Seed(time.Now().UnixNano())

	// New board
	snake.board = NewBoard(5, 15, float64(layoutWidth-5), float64(layoutHeight-5))

	// Load fonts
	tt := LoadFont("mplus-1p-regular.ttf")

	var err error
	const dpi = 72
	snake.mplusSmallFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    10,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Store important data
	snake.gridSizeX = snake.board.Width() / gridSize
	snake.gridSizeY = snake.board.Height() / gridSize
	snake.visualSize = gridSize
	snake.tickMove = 0
	snake.tickCherry = 0

	// Add the head of the snake at random position on the grid
	snake.parts = make([]*Part, 1)
	snake.parts[0] = NewPart(rand.Intn(snake.gridSizeX), rand.Intn(snake.gridSizeY), snake.board)
	currentX := snake.parts[0].X
	currentY := snake.parts[0].Y
	// Based on the position of the head, define the directions
	snake.initDirection()
	snake.applyMove()
	snake.parts = append(snake.parts, NewPart(currentX, currentY, snake.board))

	// Create an empty cherry vector
	snake.cherries = make([]*Cherry, 0)

	// Based on the position of the head, define the directions
	snake.initDirection()

	return snake
}

func (s *Snake) UpdateKeys(keys []ebiten.Key) error {
	s.ApplyDirection(keys)

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
	return len(s.parts) - 2
}

func (s *Snake) Draw(screen *ebiten.Image) {
	s.board.Draw(screen)

	scoreToDisplay := "Score: " + strconv.Itoa(len(s.parts)-2)
	placeholderScore := text.BoundString(s.mplusSmallFont, scoreToDisplay)
	text.Draw(screen, scoreToDisplay, s.mplusSmallFont, 4, placeholderScore.Dy()+2, color.White)

	for i := range s.parts {
		if i == 0 {
			s.parts[i].Draw(screen, s.direction, Head)
		} else if i == len(s.parts)-1 {
			s.parts[i].Draw(screen, s.computePartDirection(i), Tail)
		} else {
			s.parts[i].Draw(screen, s.computePartDirection(i), Body)
		}
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
	switch s.direction {
	case Up:
		s.parts[0].Y = s.parts[0].Y - 1
	case Right:
		s.parts[0].X = s.parts[0].X + 1
	case Down:
		s.parts[0].Y = s.parts[0].Y + 1
	case Left:
		s.parts[0].X = s.parts[0].X - 1
	}
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
		if float32(s.parts[0].X)/float32(s.gridSizeX) < 0.5 {
			// We are on the left side of the board, go to right
			s.direction = Right
		} else {
			// We are on the right side of the board, go to left
			s.direction = Left
		}
	} else {
		// Vertically
		if float32(s.parts[0].Y)/float32(s.gridSizeY) < 0.5 {
			// We are on the upper side of the board, go down
			s.direction = Down
		} else {
			// We are on the lower side of the board, go up
			s.direction = Up
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
		s.direction = Right
	case ebiten.KeyArrowLeft:
		s.direction = Left
	case ebiten.KeyArrowUp:
		s.direction = Up
	case ebiten.KeyArrowDown:
		s.direction = Down
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

	newTail := NewPart(lastX, lastY, s.board)
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
	x = s.parts[0].X
	y = s.parts[0].Y

	switch s.direction {
	case Up:
		y = s.parts[0].Y - 1
	case Down:
		y = s.parts[0].Y + 1
	case Left:
		x = s.parts[0].X - 1
	case Right:
		x = s.parts[0].X + 1
	}
	return
}

func (s *Snake) computePartDirection(idx int) Direction {
	if idx == 0 {
		return s.direction
	}

	if s.parts[idx].X == s.parts[idx-1].X {
		// Vertical movement
		if s.parts[idx].Y < s.parts[idx-1].Y {
			return Down
		} else {
			return Up
		}
	} else {
		// Horizontal movement
		if s.parts[idx].X < s.parts[idx-1].X {
			return Right
		} else {
			return Left
		}
	}
}
