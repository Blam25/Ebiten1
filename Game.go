package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Game struct {
	entities []entity
	player   Sprite
}

func (g *Game) Update() error {
	deltax, deltay := g.player.move()
	rectPlayer := image.Rect(int(g.player.xpos), int(g.player.ypos), int(g.player.xpos) + 30, int(g.player.ypos) + 30)
	for _, s := range g.entities {
		s.move(deltax, deltay)
		s.collide(rectPlayer)
	}
	return nil
}

var i float64 = 0

func (g *Game) Draw(screen *ebiten.Image) {

	g.player.draw(screen)
	for _, s := range g.entities {
		x, y := s.getPosition()
		if (x-g.player.xpos) > -200 && (x-g.player.xpos) < 200 && (y-g.player.ypos) > -200 && (y-g.player.ypos) < 200 {
		s.draw(screen)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
