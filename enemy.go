package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
)

type Enemy struct {
	img       *ebiten.Image
	op        ebiten.DrawImageOptions
	xpos      float64
	ypos      float64
	moveSpeed float64
	health	  int
}

func New_Enemy() Enemy {
	s := Enemy{}
	var err error
	s.img, _, err = ebitenutil.NewImageFromFile("assets/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	s.xpos = 0
	s.ypos = 0
	s.moveSpeed = 7
	s.health = 1000
	return s
}

func (s *Enemy) draw(screen *ebiten.Image) {
	s.op.GeoM.Reset()
	s.op.GeoM.Translate(s.xpos, s.ypos)
	screen.DrawImage(s.img, &s.op)
}

func (s *Enemy) moveCamera(deltaX float64, deltaY float64) {
	s.ypos = s.ypos + deltaY
	s.xpos = s.xpos + deltaX
}

func (s *Enemy) collide(rectPlayer image.Rectangle) {
	
	rectEnemy := image.Rect(int(s.xpos), int(s.ypos), int(s.xpos) + 30, int(s.ypos) + 30)
	//rectPlayer := image.Rect(int(player.xpos), int(player.ypos), int(player.xpos) + 30, int(player.ypos) + 30)
	if rectEnemy.Overlaps(rectPlayer) {
		println("yo")
		println("yo1")
	}
}

func (s *Enemy) getPosition() (float64, float64){
	return s.xpos, s.ypos
}

