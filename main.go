package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ostaquet/go-snake/snake"
	"log"
)

const layoutWidth = 320
const layoutHeight = 240
const squareSize = 10
const gridSizeX = layoutWidth / squareSize
const gridSizeY = layoutHeight / squareSize

type Game struct {
	gridSizeX, gridSizeY int
	snake                *snake.Snake
	state                GameState
	keys                 []ebiten.Key
}

const (
	GameRunning GameState = iota
	GameOver
)

type GameState int

func (g *Game) Update() error {
	// Capture keys and apply direction on Snake
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.snake.ApplyDirection(g.keys)

	// Update the state of the snake
	err := g.snake.Update()
	if errors.Is(err, snake.ErrUnauthorizedMove) {
		g.state = GameOver
	} else {
		g.state = GameRunning
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case GameRunning:
		ebitenutil.DebugPrint(screen, "Game running")
	case GameOver:
		ebitenutil.DebugPrint(screen, "Game over")
	}

	g.snake.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return layoutWidth, layoutHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Snake")
	if err := ebiten.RunGame(&Game{
		state:     GameRunning,
		gridSizeX: gridSizeX,
		gridSizeY: gridSizeY,
		snake:     snake.NewSnake(gridSizeX, gridSizeY, squareSize),
	}); err != nil {
		log.Fatal(err)
	}
}
