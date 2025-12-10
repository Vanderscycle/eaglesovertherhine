package utils

import (

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func ScreenCenterText(screen *ebiten.Image, sentence string, yOffset float64, xOffset float64, face *text.GoTextFace) (posX float64, posY float64) {
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
