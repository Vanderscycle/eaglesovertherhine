package engine

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/vanderscycle/eaglesovertherhine/assets"
)

func (g *GameEngine) startScreen(screen *ebiten.Image, margin float64) {
    if g.startImage == nil {
        return
    }

    // Get dimensions
    screenWidth := float64(screen.Bounds().Dx())
    screenHeight := float64(screen.Bounds().Dy())
    imgWidth := float64(g.startImage.Bounds().Dx())
    imgHeight := float64(g.startImage.Bounds().Dy())

    // Calculate available area after margin
    availableWidth := screenWidth - (2 * margin)
    availableHeight := screenHeight - (2 * margin)

    // Calculate scale to fit within the available area (maintain aspect ratio)
    scaleX := availableWidth / imgWidth
    scaleY := availableHeight / imgHeight
    scale := math.Min(scaleX, scaleY)

    // Calculate centered position within the available area
    scaledWidth := imgWidth * scale
    scaledHeight := imgHeight * scale

    // Center in the available area, then add margin to position
    posX := margin + (availableWidth - scaledWidth) / 2
    posY := margin + (availableHeight - scaledHeight) / 2

    // Draw image
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Scale(scale, scale)
    opts.GeoM.Translate(posX, posY)

    screen.DrawImage(g.startImage, opts)
    // write the text
    text.Draw(
        screen,
        "test0oO",
        text.NewGoXFace(assets.GameFont),
        &text.DrawOptions{
            // ColorScale: text.NewColorScale(color.White),
            // GeoM: ebiten.GeoM{}.Translate(screenWidth/2, screenHeight - margin - 30),
            // Anchor: text.AnchorCenter, // Center the text horizontally
        },
    )
}
