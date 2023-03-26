package main

import (
	_ "image/png"
	"log"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var game Game

//var op *ebiten.DrawImageOptions

func init() {
	game = Game{}
	sprite1 := New_Sprite()
	enemy1  := New_Enemy()
	enemy2  := New_Enemy() ; enemy2.xpos = 500 ; enemy2.ypos = 500
	enemy3  := New_Enemy() ; enemy3.xpos = 1000 ; enemy3.ypos = 1000
	game.entities = append(game.entities, &enemy1, &enemy2, &enemy3)
	game.player = sprite1
	//op := &ebiten.DrawImageOptions{}
	//op = &ebiten.DrawImageOptions{}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	//ebiten.SetVsyncEnabled(false)
	//ebiten.SetTPS(ebiten.SyncWithFPS)
	//ebiten.SetFPSMode(ebiten.)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

type entity interface {
	getimg() *ebiten.Image
	draw(screen *ebiten.Image)
	move(float64, float64)
	collide(image.Rectangle)
	getPosition() (float64, float64)
}
