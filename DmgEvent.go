package main

type DmgEvent struct {
	Damage int
	hitMessage string
}

func NewDmgEvent(damage int, hitMessage string) DmgEvent {
	new := DmgEvent {}
	new.Damage = damage
	new.hitMessage = hitMessage
	return new
}

func (s DmgEvent) execute(enemy *Enemy) {
	enemy.health = enemy.health - s.Damage
	print(s.hitMessage)
}