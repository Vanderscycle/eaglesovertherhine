package utils

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func ScreenCenterImage(image *ebiten.Image, screen *ebiten.Image, margin float64) (scale float64, posX float64, posY float64) {
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
