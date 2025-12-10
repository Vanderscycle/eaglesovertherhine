package engine

import (

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/vanderscycle/eaglesovertherhine/assets"
	"github.com/vanderscycle/eaglesovertherhine/engine/utils"
)

// technically this centers an image accross the whole screen
// for menus and layout I'll have to craete my own css to layout
func startScreen1(image *ebiten.Image, sentence string, screen *ebiten.Image, margin float64) {
	// Draw image
	scale, posX, posY := utils.ScreenCenterImage(image, screen, margin)
	optsImg := &ebiten.DrawImageOptions{}
	optsImg.GeoM.Scale(scale, scale)
	optsImg.GeoM.Translate(posX, posY)

	screen.DrawImage(image, optsImg)

	// write the text
	JetBrainsFace := &text.GoTextFace{
		Source:    assets.JetBrainsFaceSource,
		Direction: text.DirectionLeftToRight,
		Size:      24,
	}
	textPosX, textPosY := utils.ScreenCenterText(screen, sentence, 20.0, 1.0, JetBrainsFace)

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

func startScreen2(image *ebiten.Image, sentence string, screen *ebiten.Image, margin float64) {
	// Draw image
	scale, posX, posY := utils.ScreenCenterImage(image, screen, margin)
	optsImg := &ebiten.DrawImageOptions{}
	optsImg.GeoM.Scale(scale, scale)
	optsImg.GeoM.Translate(posX, posY)

	screen.DrawImage(image, optsImg)

	// write the text
	JetBrainsFace := &text.GoTextFace{
		Source:    assets.JetBrainsFaceSource,
		Direction: text.DirectionLeftToRight,
		Size:      24,
	}
	textPosX, textPosY := utils.ScreenCenterText(screen, sentence, 20.0, 1.0, JetBrainsFace)

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
