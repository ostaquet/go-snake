package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ostaquet/go-snake/snake"
	"log"
)

// Size of the layout/window
const layoutWidth = 320
const layoutHeight = 240

type Game struct {
	gridSizeX, gridSizeY int
	menu                 *snake.Menu
	snake                *snake.Snake
	gameOver             *snake.GameOver
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
		startGame := g.menu.UpdateKeys(g.keys)

		if startGame {
			g.state = GameStateRunning
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

		// Update the state based on pressed keys
		retry := g.gameOver.UpdateKeys(g.keys)

		if retry {
			g.snake = snake.NewSnake(layoutWidth, layoutHeight)
			g.state = GameStateRunning
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case GameStateMenu:
		g.menu.Draw(screen)
	case GameStateRunning:
		g.snake.Draw(screen)
	case GameStateOver:
		g.gameOver.Draw(screen, g.snake.Score())
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
		state:    GameStateMenu,
		menu:     snake.NewMenu(layoutWidth, layoutHeight),
		snake:    snake.NewSnake(layoutWidth, layoutHeight),
		gameOver: snake.NewGameOver(layoutWidth, layoutHeight),
	}); err != nil {
		log.Fatal(err)
	}
}
