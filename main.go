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
	menu                 *snake.Menu
	snake                *snake.Snake
	state                GameState
	keys                 []ebiten.Key
}

const (
	GameStateMenu GameState = iota
	GameStateRunning
	GameStateOver
)

type GameState int

func (g *Game) Update() error {
	// Capture keys
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	// Depending to the game state...
	switch g.state {
	case GameStateMenu:
		// Update the state of the menu based on pressed keys
		err := g.menu.UpdateKeys(g.keys)

		if err != nil {
			log.Fatal(err)
		}

	case GameStateRunning:
		// Update the state of the snake based on pressed keys
		err := g.snake.UpdateKeys(g.keys)

		// If there is an unauthorized move, it is a game over
		if errors.Is(err, snake.ErrUnauthorizedMove) {
			g.state = GameStateOver
		} else {
			if err != nil {
				log.Fatal(err)
			}
			g.state = GameStateRunning
		}
	case GameStateOver:
		// Stay in Gome Over state
		g.state = GameStateOver
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	switch g.state {
	case GameStateMenu:

	case GameStateRunning:
		g.snake.Draw(screen)
		ebitenutil.DebugPrint(screen, "Score "+strconv.Itoa(g.snake.Score()))
	case GameStateOver:
		g.snake.Draw(screen)
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
		state: GameStateRunning,
		menu:  snake.NewMenu(layoutWidth, layoutHeight),
		snake: snake.NewSnake(layoutWidth, layoutHeight),
	}); err != nil {
		log.Fatal(err)
	}
}
