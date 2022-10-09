package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ostaquet/go-snake/snake"
	"log"
	"strconv"
)

// Size of the layout/window
const layoutWidth = 320
const layoutHeight = 240

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
	// Depending to the game state...
	switch g.state {
	case GameRunning:
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
	case GameOver:
		g.state = GameOver
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.snake.Draw(screen)

	switch g.state {
	case GameRunning:
		ebitenutil.DebugPrint(screen, "Score "+strconv.Itoa(g.snake.Score()))
	case GameOver:
		ebitenutil.DebugPrint(screen, "Game over\nScore "+strconv.Itoa(g.snake.Score()))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return layoutWidth, layoutHeight
}

func main() {
	// Prepare the screen
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Snake")

	if err := ebiten.RunGame(&Game{
		state: GameRunning,
		snake: snake.NewSnake(layoutWidth, layoutHeight),
	}); err != nil {
		log.Fatal(err)
	}
}
