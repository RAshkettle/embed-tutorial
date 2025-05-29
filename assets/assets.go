package assets

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

var PlayerSpriteImage = loadImage("images/red_character.png")

func loadImage(filePath string) *ebiten.Image {
	data, err := assets.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	ebitenImg := ebiten.NewImageFromImage(img)
	return ebitenImg
}
