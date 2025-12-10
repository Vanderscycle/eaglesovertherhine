package engine

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/vanderscycle/eaglesovertherhine/assets"
)

func screenCenterImage(image *ebiten.Image, screen *ebiten.Image, margin float64) (scale float64, posX float64, posY float64) {
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
	scale = math.Min(scaleX, scaleY)

	// Calculate centered position within the available area
	scaledWidth := imgWidth * scale
	scaledHeight := imgHeight * scale

	// Center in the available area, then add margin to position
	posX = margin + (availableWidth-scaledWidth)/2
	posY = margin + (availableHeight-scaledHeight)/2
	return scale, posX, posY
}

func screenCenterText(screen *ebiten.Image, sentence string, yOffset float64, xOffset float64, face *text.GoTextFace) (posX float64, posY float64) {
	if sentence == "" {
		return
	}
	screenWidthx := float64(screen.Bounds().Dx())
	screenHeighty := float64(screen.Bounds().Dy())

	wordWidth, wordHeight := text.Measure(sentence, face, 0) //w,h
	posX = float64(screenWidthx-wordWidth-xOffset) / 2
	posY = float64(screenHeighty - wordHeight - yOffset)
	return posX, posY
}

// technically this centers an image accross the whole screen
// for menus and layout I'll have to craete my own css to layout
func startScreen(image *ebiten.Image, screen *ebiten.Image, margin float64) {
	// Draw image
	scale, posX, posY := screenCenterImage(image, screen, margin)
	optsImg := &ebiten.DrawImageOptions{}
	optsImg.GeoM.Scale(scale, scale)
	optsImg.GeoM.Translate(posX, posY)

	screen.DrawImage(image, optsImg)

	// write the text

	sentence := "Ken Thompson"

	JetBrainsFace := &text.GoTextFace{
		Source:    assets.JetBrainsFaceSource,
		Direction: text.DirectionLeftToRight,
		Size:      24,
	}
	textPosX, textPosY := screenCenterText(screen, sentence, 40.0, 1.0, JetBrainsFace)

	optsTxt := &text.DrawOptions{}
	optsTxt.ColorScale.Scale(1.0, 0.4, 0.7, 1.0) // pink (R,G,B,A)
	optsTxt.GeoM.Translate(textPosX, textPosY)

	text.Draw(
		screen,
		sentence,
		JetBrainsFace,
		optsTxt,
	)
}
