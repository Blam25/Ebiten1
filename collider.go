package main

/*
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"*/
import	"image"

type Collider struct {
	enemy *Enemy
	events []event
	Velocity float64
}

func NewCollider(game *Game, enemy *Enemy, events ...event) *Collider{
	new := Collider{}
	new.enemy = enemy
	new.events = events
	game.colliders = append(game.colliders, new)
	return &new
}

func (s *Collider) collide(rectPlayer image.Rectangle) {
	
	rectEnemy := image.Rect(int(s.enemy.xpos), int(s.enemy.ypos), int(s.enemy.xpos) + 30, int(s.enemy.ypos) + 30)
	//rectPlayer := image.Rect(int(player.xpos), int(player.ypos), int(player.xpos) + 30, int(player.ypos) + 30)
	if rectEnemy.Overlaps(rectPlayer) {
		for _,z := range s.events {
			z.execute(s.enemy)
		}
	}
}


