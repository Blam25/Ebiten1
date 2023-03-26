package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var game Game

//var op *ebiten.DrawImageOptions

func init() {
	game = NewGame()
	sprite1 := New_Sprite()
	enemy1 := New_Enemy()
	enemy2 := New_Enemy() ; enemy2.xpos = 500 ; enemy2.ypos = 500 ; NewCollider(&game, &enemy2, NewDmgEvent(60, "Im hit"), NewDmgEvent(60, "Yes really"))
	enemy3 := New_Enemy() ; enemy3.xpos = 1000 ; enemy3.ypos = 1000 ;
	game.entities = append(game.entities, &enemy1, &enemy2, &enemy3)
	game.collideables = append(game.collideables, &enemy1, &enemy2, &enemy3)
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
	draw(screen *ebiten.Image)
	moveCamera(float64, float64)
	getPosition() (float64, float64)
}

type collide interface {
	collide(image.Rectangle)
}

type event interface {
	execute(enemy *Enemy)
}
