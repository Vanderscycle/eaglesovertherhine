package engine

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/vanderscycle/eaglesovertherhine/assets"
)

// technically this centers an image accross the whole screen
// for menus and layout I'll have to craete my own css to layout
func startScreen(image *ebiten.Image, screen *ebiten.Image, margin float64) {
	// TODO create an error
	if image == nil {
		return
	}

	// Get dimensions
	screenWidth := float64(screen.Bounds().Dx())
	screenHeight := float64(screen.Bounds().Dy())
	imgWidth := float64(image.Bounds().Dx())
	imgHeight := float64(image.Bounds().Dy())

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
	posX := margin + (availableWidth-scaledWidth)/2
	posY := margin + (availableHeight-scaledHeight)/2

	// Draw image
	optsImg := &ebiten.DrawImageOptions{}
	optsImg.GeoM.Scale(scale, scale)
	optsImg.GeoM.Translate(posX, posY)

	screen.DrawImage(image, optsImg)
	//---

		face := &text.GoTextFace{
			Source:    assets.JetBrainsFaceSource,
			Direction: text.DirectionLeftToRight,
			Size:      24,
		}
w, _ := text.Measure("Ken Thompson", face, 0) //w,h
	// write the text
	optsTxt := &text.DrawOptions{}
	optsTxt.ColorScale.Scale(1.0, 0.4, 0.7, 1.0) // pink (R,G,B,A)
	optsTxt.GeoM.Translate(float64(screenWidth-w)/2, float64(screenHeight-margin-40))

	//optsTxt.GeoM :=ebiten.GeoM{}.Translate(screenWidth/2, screenHeight - margin - 30)

	text.Draw(
		screen,
		"Ken Thompson",
		face,
		optsTxt,
	)
}
