package engine

import (
	//"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameEngine struct{}

func (g *GameEngine) Update() error {
	return nil
}

func (g *GameEngine) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

}

func (g *GameEngine) Layout(outsideWidth, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() *GameEngine {
	g := &GameEngine{}
	ebiten.SetWindowTitle("Eagles Over the Rhine")
	// eventually you'd want a menu to change the size
	ebiten.SetWindowSize( DefaultOptionMenu.width,DefaultOptionMenu.height)
	return g
}
