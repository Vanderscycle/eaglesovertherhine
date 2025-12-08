package engine

import "github.com/hajimehoshi/ebiten"

type OptionMenu struct{
	height int
	width int
}
var DefaultOptionMenu = OptionMenu{
	height: 1024,
	width: 768,
}

func (g *GameEngine) Option(optionMenu OptionMenu) {
	ebiten.SetWindowSize( optionMenu.width,optionMenu.height)
	g.Layout(optionMenu.width, optionMenu.height)
}
