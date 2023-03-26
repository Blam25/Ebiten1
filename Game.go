package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	entities       []entity
	collideables	   []collide
	player         Sprite
	renderDistance float64
	colliders		[]Collider
}

func NewGame() Game {
	new := Game{}
	new.renderDistance = 500
	return new
}

func (g *Game) Update() error {
	deltax, deltay := g.player.move()
	for _, s := range g.entities {
		s.moveCamera(deltax, deltay)
	}
	rectPlayer := image.Rect(int(g.player.xpos), int(g.player.ypos), int(g.player.xpos)+30, int(g.player.ypos)+30)
	/*for _, s := range g.collideables {
		s.collide(rectPlayer)
	}*/
	for _, s := range g.colliders {
		s.collide(rectPlayer)
	}
	return nil
}

var i float64 = 0

func (g *Game) Draw(screen *ebiten.Image) {

	g.player.draw(screen)
	for _, s := range g.entities {
		x, y := s.getPosition()
		if (x-g.player.xpos) > -g.renderDistance && (x-g.player.xpos) < g.renderDistance && (y-g.player.ypos) > -g.renderDistance && (y-g.player.ypos) < g.renderDistance {
			s.draw(screen)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
