package main

import (
	"embed_tutorial/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *ebiten.Image
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(100.0, 30.0)
	screen.DrawImage(g.player, &op)
}

func (g *Game) Layout(outerScreenWidth, outerScreenHeight int) (ScreenWidth, ScreenHeight int) {
	return 480, 320
}

func (g *Game) Update() error {
	return nil
}

func main() {
	g := &Game{}
	g.player = assets.PlayerSpriteImage

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
