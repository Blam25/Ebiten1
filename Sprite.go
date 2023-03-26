package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	img       *ebiten.Image
	op        ebiten.DrawImageOptions
	xpos      float64
	ypos      float64
	moveSpeed float64
}

func New_Sprite() Sprite {
	s := Sprite{}
	var err error
	s.img, _, err = ebitenutil.NewImageFromFile("assets/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	s.xpos = 150
	s.ypos = 150
	s.moveSpeed = 7
	return s
}

func (s *Sprite) draw(screen *ebiten.Image) {
	s.op.GeoM.Reset()
	s.op.GeoM.Translate(s.xpos, s.ypos)
	screen.DrawImage(s.img, &s.op)
}


func (s *Sprite) move() (float64, float64) {
	var x float64
	var y float64
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		//s.ypos = s.ypos - s.moveSpeed
		y = s.moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		//s.ypos = s.ypos + s.moveSpeed
		y = -s.moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		//s.xpos = s.xpos + s.moveSpeed
		x = -s.moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		//s.xpos = s.xpos - s.moveSpeed
		x = s.moveSpeed
	}
	return x, y
}
