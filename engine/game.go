package engine

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/vanderscycle/eaglesovertherhine/assets"
)

type GameEngine struct {
	gameState     int
	introStartTime time.Time
}

const (
	stateIntro = iota
	stateGame
	stateMenu
)
func (g *GameEngine) Update() error {
	switch g.gameState {
	case stateIntro:
		// Check if intro time has elapsed (3 seconds = 1s wait + 2s display)
		if time.Since(g.introStartTime) >= 5*time.Second {
			g.gameState = stateMenu // or stateGame
		}
	case stateGame:
		// Game logic here
	case stateMenu:
		// Menu logic here
	}
	return nil
}

func (g *GameEngine) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	// Clear screen with black
	//screen.Fill(black) // Define black color or use ebiten.ColorScale{}

	switch g.gameState {
	case stateIntro:
		// Only show start image after 1 second
		if time.Since(g.introStartTime) >= 1*time.Second {
  startScreen(assets.StartSprites,screen, 25)
		}
		// First second is just black screen
	case stateGame:
		// Draw game
	case stateMenu:
		// Draw menu
	}

}

func (g *GameEngine) Layout(outsideWidth, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() *GameEngine {
	g := &GameEngine{
		gameState:     stateIntro,
		introStartTime: time.Now(),

	}
	ebiten.SetWindowTitle("Eagles Over the Rhine")
	// eventually you'd want a menu to change the size
	ebiten.SetWindowSize( DefaultOptionMenu.width,DefaultOptionMenu.height)
	return g
}
