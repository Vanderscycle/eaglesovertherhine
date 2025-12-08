package engine

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vanderscycle/eaglesovertherhine/assets"
)

func Start(g *GameEngine, screen *ebiten.Image) {
	time.Sleep(1 * time.Second)
	start := assets.StartSprites
	screen.DrawImage(start, nil)
	time.Sleep(2 * time.Second)
}
